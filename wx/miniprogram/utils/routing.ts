export namespace routing {
    // driving
    export interface DriveOpts {
        trip_id: string
    }

    export function driving(opt: DriveOpts) {
        return `/pages/driving/driving?trip_id=${opt.trip_id}`
    }

    // lock
    export interface LockOpts {
        car_id: string
    }

    export function lock(opt: LockOpts) {
        return `/pages/lock/lock?car_id=${opt.car_id}`
    }

    // register
    export interface RegisterOpts {
        redirect?: string
    }

    export interface RegisterParams {
        redirectURL: string
    }

    export function register(opt?: RegisterParams) {
        const page = '/pages/register/register'
        if (!opt) {
            return page
        }
        return `${page}?redirect=${encodeURIComponent(opt.redirectURL)}`
    }

    // mytrips
    export function mytrips() {
        return '/pages/mytrips/mytrips'
    }

    // index
    export function index() {
        return '/pages/index/index'
    }
}
