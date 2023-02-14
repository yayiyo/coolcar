import {Identity, Profile} from "./proto_gen/rental/rental";
import {CoolCar} from "./request";

export namespace ProfileService {
    export function getProfile(): Promise<Profile> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'GET',
            path: '/v1/profile',
            resMarshal: Profile.fromJson,
        })
    }

    export function submitProfile(req: Identity): Promise<Profile> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'POST',
            path: '/v1/profile',
            data: req,
            resMarshal: Profile.fromJson,
        })
    }

    export function clearProfile(): Promise<Profile> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'DELETE',
            path: '/v1/profile',
            resMarshal: Profile.fromJson,
        })
    }
}