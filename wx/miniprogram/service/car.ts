import {CarEntity} from "./proto_gen/car/car";
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
}