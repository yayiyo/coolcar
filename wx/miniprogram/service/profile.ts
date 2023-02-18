import {
    ClearProfilePhotoResponse,
    CreateProfilePhotoResponse,
    GetProfilePhotoResponse,
    Identity,
    Profile
} from "./proto_gen/rental/rental";
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

    export function getProfilePhoto(): Promise<GetProfilePhotoResponse> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'GET',
            path: '/v1/profile/photo',
            resMarshal: GetProfilePhotoResponse.fromJson,
        })
    }

    export function createProfilePhoto(): Promise<CreateProfilePhotoResponse> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'POST',
            path: '/v1/profile/photo',
            resMarshal: CreateProfilePhotoResponse.fromJson,
        })
    }

    export function completeProfilePhoto(): Promise<Identity> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'POST',
            path: '/v1/profile/photo/complete',
            resMarshal: Identity.fromJson,
        })
    }

    export function clearProfilePhoto(): Promise<ClearProfilePhotoResponse> {
        return CoolCar.sendRequestWithAuthRetry({
            method: 'DELETE',
            path: '/v1/profile/photo',
            resMarshal: ClearProfilePhotoResponse.fromJson,
        })
    }
}