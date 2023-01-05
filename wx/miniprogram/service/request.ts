import {LoginRequest, LoginResponse} from "./proto_gen/auth/auth";
import camelcaseKeys from "camelcase-keys";

export namespace CoolCar {
    const serverAddr = 'http://localhost:8080'
    const AUTH_ERROR = 'AUTH_ERROR'
    const authData = {
        token: '',
        expireMs: 0 // 过期时间：毫秒
    }

    export interface RequestOption<REQ, RES> {
        method: 'GET' | 'POST' | 'PUT' | 'DELETE',
        path: string,
        data: REQ,
        resMarshal: (json: string) => RES,
    }

    export interface AuthOption {
        attachAuthHeader: boolean
        retryOnAuthError: boolean
    }

    export async function sendRequestWithAuthRetry<REQ, RES>(o: RequestOption<REQ, RES>, a?: AuthOption): Promise<RES> {
        const authOpt = a || {
            attachAuthHeader: true,
            retryOnAuthError: true,
        }
        try {
            await login()
            return sendRequest(o, authOpt)
        } catch (err) {
            if (err === AUTH_ERROR && authOpt.retryOnAuthError) {
                authData.token = ''
                authData.expireMs = 0
                return sendRequestWithAuthRetry(o, {
                    attachAuthHeader: authOpt.attachAuthHeader,
                    retryOnAuthError: false,
                })
            } else {
                throw err
            }
        }
    }

    export async function login() {
        // 如果token有效 无需再次登陆
        if (authData.token && authData.expireMs >= Date.now()) {
            return
        }
        const wxResp = await wxLogin()
        const reqTimeMs = Date.now() // 毫秒
        const resp = await sendRequest<LoginRequest, LoginResponse>({
            method: 'POST',
            path: '/v1/auth/login',
            data: {
                code: wxResp.code,
            },
            resMarshal: LoginResponse.fromJson,
        }, {
            attachAuthHeader: false,
            retryOnAuthError: false,
        })
        console.log('--------------------login--------------------')
        console.log(resp)
        authData.token = resp.accessToken
        authData.expireMs = reqTimeMs + resp.expiresIn * 1000
    }

    function sendRequest<REQ, RES>(o: RequestOption<REQ, RES>, a: AuthOption): Promise<RES> {
        return new Promise<RES>((resolve, reject) => {
            const header: Record<string, any> = {}
            if (a.attachAuthHeader) {
                if (authData.token && authData.expireMs >= Date.now()) {
                    header.authorization = `Bearer ${authData.token}`
                } else {
                    reject(AUTH_ERROR)
                    return
                }
            }
            wx.request({
                url: serverAddr + o.path,
                method: o.method,
                data: o.data as object,
                header: header,
                success: res => {
                    if (res.statusCode === 401) {
                        reject(AUTH_ERROR)
                    } else if (res.statusCode >= 400) {
                        reject(res)
                    } else {
                        console.log(res.data)
                        const toCamelObj = camelcaseKeys(
                            res.data as object,
                            {
                                deep: true
                            })
                        const obj = JSON.parse(JSON.stringify(toCamelObj)) as RES
                        console.log(obj)
                        resolve(obj)
                    }

                },
                fail: reject,
            })
        })
    }

    function wxLogin(): Promise<WechatMiniprogram.LoginSuccessCallbackResult> {
        return new Promise((resolve, reject) => {
            wx.login({
                success: result => {
                    console.log(result)
                    resolve(result)
                },
                fail: res => {
                    console.error(res)
                    reject(res)
                },
            })
        })
    }
}