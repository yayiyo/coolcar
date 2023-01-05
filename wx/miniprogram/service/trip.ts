import {CreateTripRequest, CreateTripResponse} from "./proto_gen/rental/rental";
import {CoolCar} from "./request";

export namespace TripService {
    export function createTrip(req: CreateTripRequest): Promise<CreateTripResponse> {
        return CoolCar.sendRequestWithAuthRetry({
            path: '/v1/trip',
            method: 'POST',
            data: req,
            resMarshal: CreateTripResponse.fromJson,
        })
    }
}