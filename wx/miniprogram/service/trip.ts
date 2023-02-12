import {
    CreateTripRequest,
    GetTripRequest,
    GetTripsResponse,
    Trip,
    TripEntity,
    TripStatus
} from "./proto_gen/rental/rental";
import {CoolCar} from "./request";

export namespace TripService {
    export function createTrip(req: CreateTripRequest): Promise<TripEntity> {
        return CoolCar.sendRequestWithAuthRetry({
            path: '/v1/trips',
            method: 'POST',
            data: req,
            resMarshal: TripEntity.fromJson,
        })
    }

    export function getTrip(id :string): Promise<Trip> {
        return CoolCar.sendRequestWithAuthRetry({
            path: `/v1/trips/${encodeURIComponent(id)}`,
            method: 'GET',
            resMarshal: Trip.fromJson,
        })
    }

    export function getTrips(s?: TripStatus): Promise<GetTripsResponse> {
        let path = '/v1/trips'
        if (s) {
            path += `?status=${s}`
        }
        return CoolCar.sendRequestWithAuthRetry({
            path,
            method: 'GET',
            resMarshal: GetTripsResponse.fromJson,
        })
    }
}