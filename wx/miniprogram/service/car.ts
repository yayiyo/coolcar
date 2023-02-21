import {Car, CarEntity} from "./proto_gen/car/car";
import {CoolCar} from "./request";
import camelcaseKeys from "camelcase-keys";

export namespace CarService {

    export function subscribe(onMsg: (car: CarEntity) => void) {
        const socket = wx.connectSocket({
            url: CoolCar.wsAddr + '/ws',
        })

        socket.onMessage(msg => {
            const obj = JSON.parse(msg.data as string)
            onMsg(CarEntity.fromJson(camelcaseKeys(obj, {
                deep: true,
            })))
        })
        return socket
    }

    export function getCar(id: string): Promise<Car> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'GET',
            path: `/v1/car/${encodeURIComponent(id)}`,
            resMarshal: Car.fromJson,
        })
    }
}