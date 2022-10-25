import * as $protobuf from "protobufjs";
/** Namespace coolcar. */
export namespace coolcar {

    /** Properties of a Location. */
    interface ILocation {

        /** Location latitude */
        latitude?: (number|null);

        /** Location longitude */
        longitude?: (number|null);
    }

    /** Represents a Location. */
    class Location implements ILocation {

        /**
         * Constructs a new Location.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ILocation);

        /** Location latitude. */
        public latitude: number;

        /** Location longitude. */
        public longitude: number;

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Location
         */
        public static fromObject(object: { [k: string]: any }): coolcar.Location;

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @param message Location
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.Location, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Location to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a LocationStatus. */
    interface ILocationStatus {

        /** LocationStatus location */
        location?: (coolcar.ILocation|null);

        /** LocationStatus feeCent */
        feeCent?: (number|null);

        /** LocationStatus kmDriven */
        kmDriven?: (number|null);

        /** LocationStatus poiName */
        poiName?: (string|null);

        /** LocationStatus timestampSec */
        timestampSec?: (number|null);
    }

    /** Represents a LocationStatus. */
    class LocationStatus implements ILocationStatus {

        /**
         * Constructs a new LocationStatus.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ILocationStatus);

        /** LocationStatus location. */
        public location?: (coolcar.ILocation|null);

        /** LocationStatus feeCent. */
        public feeCent: number;

        /** LocationStatus kmDriven. */
        public kmDriven: number;

        /** LocationStatus poiName. */
        public poiName: string;

        /** LocationStatus timestampSec. */
        public timestampSec: number;

        /**
         * Creates a LocationStatus message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns LocationStatus
         */
        public static fromObject(object: { [k: string]: any }): coolcar.LocationStatus;

        /**
         * Creates a plain object from a LocationStatus message. Also converts values to other types if specified.
         * @param message LocationStatus
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.LocationStatus, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this LocationStatus to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** TripStatus enum. */
    enum TripStatus {
        TS_NOT_SPECIFIED = 0,
        IN_PROGRESS = 1,
        FINISHED = 2
    }

    /** Properties of a TripEntity. */
    interface ITripEntity {

        /** TripEntity id */
        id?: (string|null);

        /** TripEntity trip */
        trip?: (coolcar.ITrip|null);
    }

    /** Represents a TripEntity. */
    class TripEntity implements ITripEntity {

        /**
         * Constructs a new TripEntity.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ITripEntity);

        /** TripEntity id. */
        public id: string;

        /** TripEntity trip. */
        public trip?: (coolcar.ITrip|null);

        /**
         * Creates a TripEntity message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns TripEntity
         */
        public static fromObject(object: { [k: string]: any }): coolcar.TripEntity;

        /**
         * Creates a plain object from a TripEntity message. Also converts values to other types if specified.
         * @param message TripEntity
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.TripEntity, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this TripEntity to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a Trip. */
    interface ITrip {

        /** Trip accountId */
        accountId?: (string|null);

        /** Trip carId */
        carId?: (string|null);

        /** Trip start */
        start?: (coolcar.ILocationStatus|null);

        /** Trip current */
        current?: (coolcar.ILocationStatus|null);

        /** Trip end */
        end?: (coolcar.ILocationStatus|null);

        /** Trip status */
        status?: (coolcar.TripStatus|null);

        /** Trip identityId */
        identityId?: (string|null);
    }

    /** Represents a Trip. */
    class Trip implements ITrip {

        /**
         * Constructs a new Trip.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ITrip);

        /** Trip accountId. */
        public accountId: string;

        /** Trip carId. */
        public carId: string;

        /** Trip start. */
        public start?: (coolcar.ILocationStatus|null);

        /** Trip current. */
        public current?: (coolcar.ILocationStatus|null);

        /** Trip end. */
        public end?: (coolcar.ILocationStatus|null);

        /** Trip status. */
        public status: coolcar.TripStatus;

        /** Trip identityId. */
        public identityId: string;

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Trip
         */
        public static fromObject(object: { [k: string]: any }): coolcar.Trip;

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @param message Trip
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.Trip, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Trip to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a CreateTripRequest. */
    interface ICreateTripRequest {

        /** CreateTripRequest start */
        start?: (coolcar.ILocation|null);

        /** CreateTripRequest carId */
        carId?: (string|null);

        /** CreateTripRequest avatarUrl */
        avatarUrl?: (string|null);
    }

    /** Represents a CreateTripRequest. */
    class CreateTripRequest implements ICreateTripRequest {

        /**
         * Constructs a new CreateTripRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ICreateTripRequest);

        /** CreateTripRequest start. */
        public start?: (coolcar.ILocation|null);

        /** CreateTripRequest carId. */
        public carId: string;

        /** CreateTripRequest avatarUrl. */
        public avatarUrl: string;

        /**
         * Creates a CreateTripRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns CreateTripRequest
         */
        public static fromObject(object: { [k: string]: any }): coolcar.CreateTripRequest;

        /**
         * Creates a plain object from a CreateTripRequest message. Also converts values to other types if specified.
         * @param message CreateTripRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.CreateTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this CreateTripRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a GetTripRequest. */
    interface IGetTripRequest {

        /** GetTripRequest id */
        id?: (string|null);
    }

    /** Represents a GetTripRequest. */
    class GetTripRequest implements IGetTripRequest {

        /**
         * Constructs a new GetTripRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.IGetTripRequest);

        /** GetTripRequest id. */
        public id: string;

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripRequest
         */
        public static fromObject(object: { [k: string]: any }): coolcar.GetTripRequest;

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @param message GetTripRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.GetTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a GetTripsRequest. */
    interface IGetTripsRequest {

        /** GetTripsRequest status */
        status?: (coolcar.TripStatus|null);
    }

    /** Represents a GetTripsRequest. */
    class GetTripsRequest implements IGetTripsRequest {

        /**
         * Constructs a new GetTripsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.IGetTripsRequest);

        /** GetTripsRequest status. */
        public status: coolcar.TripStatus;

        /**
         * Creates a GetTripsRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripsRequest
         */
        public static fromObject(object: { [k: string]: any }): coolcar.GetTripsRequest;

        /**
         * Creates a plain object from a GetTripsRequest message. Also converts values to other types if specified.
         * @param message GetTripsRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.GetTripsRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripsRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a GetTripsResponse. */
    interface IGetTripsResponse {

        /** GetTripsResponse trips */
        trips?: (coolcar.ITripEntity[]|null);
    }

    /** Represents a GetTripsResponse. */
    class GetTripsResponse implements IGetTripsResponse {

        /**
         * Constructs a new GetTripsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.IGetTripsResponse);

        /** GetTripsResponse trips. */
        public trips: coolcar.ITripEntity[];

        /**
         * Creates a GetTripsResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripsResponse
         */
        public static fromObject(object: { [k: string]: any }): coolcar.GetTripsResponse;

        /**
         * Creates a plain object from a GetTripsResponse message. Also converts values to other types if specified.
         * @param message GetTripsResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.GetTripsResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripsResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of an UpdateTripRequest. */
    interface IUpdateTripRequest {

        /** UpdateTripRequest id */
        id?: (string|null);

        /** UpdateTripRequest current */
        current?: (coolcar.ILocation|null);

        /** UpdateTripRequest endTrip */
        endTrip?: (boolean|null);
    }

    /** Represents an UpdateTripRequest. */
    class UpdateTripRequest implements IUpdateTripRequest {

        /**
         * Constructs a new UpdateTripRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.IUpdateTripRequest);

        /** UpdateTripRequest id. */
        public id: string;

        /** UpdateTripRequest current. */
        public current?: (coolcar.ILocation|null);

        /** UpdateTripRequest endTrip. */
        public endTrip: boolean;

        /**
         * Creates an UpdateTripRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns UpdateTripRequest
         */
        public static fromObject(object: { [k: string]: any }): coolcar.UpdateTripRequest;

        /**
         * Creates a plain object from an UpdateTripRequest message. Also converts values to other types if specified.
         * @param message UpdateTripRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.UpdateTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this UpdateTripRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Represents a TripService */
    class TripService extends $protobuf.rpc.Service {

        /**
         * Constructs a new TripService service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls CreateTrip.
         * @param request CreateTripRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and TripEntity
         */
        public createTrip(request: coolcar.ICreateTripRequest, callback: coolcar.TripService.CreateTripCallback): void;

        /**
         * Calls CreateTrip.
         * @param request CreateTripRequest message or plain object
         * @returns Promise
         */
        public createTrip(request: coolcar.ICreateTripRequest): Promise<coolcar.TripEntity>;

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and Trip
         */
        public getTrip(request: coolcar.IGetTripRequest, callback: coolcar.TripService.GetTripCallback): void;

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @returns Promise
         */
        public getTrip(request: coolcar.IGetTripRequest): Promise<coolcar.Trip>;

        /**
         * Calls GetTrips.
         * @param request GetTripsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and GetTripsResponse
         */
        public getTrips(request: coolcar.IGetTripsRequest, callback: coolcar.TripService.GetTripsCallback): void;

        /**
         * Calls GetTrips.
         * @param request GetTripsRequest message or plain object
         * @returns Promise
         */
        public getTrips(request: coolcar.IGetTripsRequest): Promise<coolcar.GetTripsResponse>;

        /**
         * Calls UpdateTrip.
         * @param request UpdateTripRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and Trip
         */
        public updateTrip(request: coolcar.IUpdateTripRequest, callback: coolcar.TripService.UpdateTripCallback): void;

        /**
         * Calls UpdateTrip.
         * @param request UpdateTripRequest message or plain object
         * @returns Promise
         */
        public updateTrip(request: coolcar.IUpdateTripRequest): Promise<coolcar.Trip>;
    }

    namespace TripService {

        /**
         * Callback as used by {@link coolcar.TripService#createTrip}.
         * @param error Error, if any
         * @param [response] TripEntity
         */
        type CreateTripCallback = (error: (Error|null), response?: coolcar.TripEntity) => void;

        /**
         * Callback as used by {@link coolcar.TripService#getTrip}.
         * @param error Error, if any
         * @param [response] Trip
         */
        type GetTripCallback = (error: (Error|null), response?: coolcar.Trip) => void;

        /**
         * Callback as used by {@link coolcar.TripService#getTrips}.
         * @param error Error, if any
         * @param [response] GetTripsResponse
         */
        type GetTripsCallback = (error: (Error|null), response?: coolcar.GetTripsResponse) => void;

        /**
         * Callback as used by {@link coolcar.TripService#updateTrip}.
         * @param error Error, if any
         * @param [response] Trip
         */
        type UpdateTripCallback = (error: (Error|null), response?: coolcar.Trip) => void;
    }
}
