import requests
from operator import itemgetter

def RetrieveAirports():
    URL = "https://open-atms.airlab.aero/api/v1/airac/airports"

    response = requests.get(URL, headers = {
        "api-key": "lgBaEkJ1TLQrwFDhtwe2mqLWIgoiyxue9kmrNkvOKpdjfhyXHIcdw7MNLmTLopH6",
        'Content-Type': 'application/json',
        })

    if response.status_code == 200:
        return response.json()

def RetrieveSID(icao):
    URL = "https://open-atms.airlab.aero/api/v1/airac/sids/airport/"

    response = requests.get(f"{URL}{icao}", headers = {
        "api-key": "lgBaEkJ1TLQrwFDhtwe2mqLWIgoiyxue9kmrNkvOKpdjfhyXHIcdw7MNLmTLopH6",
        'Content-Type': 'application/json',
        })

    if response.status_code == 200:
        return response.json()

def RetrieveSTARs(icao):
    URL = "https://open-atms.airlab.aero/api/v1/airac/stars/airport/"

    response = requests.get(f"{URL}{icao}", headers = {
        "api-key": "lgBaEkJ1TLQrwFDhtwe2mqLWIgoiyxue9kmrNkvOKpdjfhyXHIcdw7MNLmTLopH6",
        'Content-Type': 'application/json',
        })

    if response.status_code == 200:
        return response.json()

def CalculateTopWaypoints(SID_list):
    sorted_SID_list = sorted(SID_list, key=itemgetter('count'), reverse=True)
    return sorted_SID_list[:2]

def Calculate(airport, TYPEs, TYPEs_airport_list):
    TYPEs_list = list()
    for TYPE in TYPEs:
        TYPEs_list.append(dict(name=TYPE['name'], count=len(TYPE['waypoints'])))

    top_waypoints = CalculateTopWaypoints(TYPEs_list)
    TYPEs_airport_list.append(dict(airport=airport['icao'], topWaypoints=top_waypoints))
    return TYPEs_airport_list

def RetrieveSTARsSIDs():
    airports = RetrieveAirports()

    SIDs_airport_list = list()
    STARs_airport_list = list()
    for airport in airports:
        STARs = RetrieveSTARs(airport['icao'])
        SIDs = RetrieveSID(airport['icao'])
        SIDs_airport_list = Calculate(airport, SIDs, SIDs_airport_list)
        STARs_airport_list = Calculate(airport, STARs, STARs_airport_list)

    return dict(SIDs_airport_list=SIDs_airport_list, STARs_airport_list=STARs_airport_list)