import { Coolcar } from "../utils/request";
import { rental } from "./proto_gen/rental/rental_pb";

export namespace TripService {
    export function CreateTrip(tripEntity:rental.v1.ICreateTripRequest):Promise<rental.v1.ITripEntity> {
        return Coolcar.sendRequestWithRetry({
            data:tripEntity,
            path:"/v1/trip",
            method:"POST",
            anyfromObjectFunc:rental.v1.TripEntity.fromObject,
           },{
            needAttachTokenHeader: true,
           },{
            needRetry: true,
            times: 0
           })
    }


    export function GetTrip(id:string):Promise<rental.v1.ITrip> {
        return Coolcar.sendRequestWithRetry({
            method:'GET',
            path:`/v1/trip/${encodeURIComponent(id)}`,
            anyfromObjectFunc:rental.v1.Trip.fromObject
        },{
            needAttachTokenHeader: true,
           },{
            needRetry: true,
            times: 0
           })
    }

    export function GetTrips(s?:rental.v1.TripStatus):Promise<rental.v1.IGetTripsResponse> {
        let path = '/v1/trips'
        if (s) {
            path += `?status=${s}`
        }
        return Coolcar.sendRequestWithRetry({
            method:'GET',
            path,
            anyfromObjectFunc:rental.v1.GetTripsResponse.fromObject
        },{
            needAttachTokenHeader: true,
           },{
            needRetry: true,
            times: 0
           })
    }


    export function updateTripPos(id:string,loc?:rental.v1.ILocation) {
        return updateTrip({
            id,
            current:loc,
        })
    }


    export function finishTrip(id:string,current: rental.v1.ILocation) {
        updateTrip({
            id,
            endTrip:true,
            current
        } as rental.v1.UpdateTripsRequest)
    }

    export function updateTrip(r: rental.v1.IUpdateTripsRequest):Promise<rental.v1.ITrip>{
        if (!r.id) {
            return Promise.reject("must sperify id")
        }
        return Coolcar.sendRequestWithRetry({
            anyfromObjectFunc: rental.v1.Trip.fromObject,
            data:r,
            path:`/v1/trip/${encodeURIComponent(r.id)}`,
            method:"PUT"
        },{
            needAttachTokenHeader: true,
        },{
            times:0,
            needRetry:true
        })
    }
}