import datetime
import requests

BASE_URL = "http://localhost:8080"

def login(username, password):
    print(f"Logging in as {username}")
    response = requests.post(f"{BASE_URL}/login", json={"username": username, "password": password})
    if response.status_code == 200:
        token = response.json()["token"]
        print(f"- Login successful. Token: {token}")
        return token
    else:
        print(f"- Login failed: {response.text}")
        return None

def create_user(username, password, role, token):
    print(f"Creating {username}...")
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.post(f"{BASE_URL}/{role}/users", json={"username": username, "password": password}, headers=headers)
    if response.status_code == 200:
        print(f"- User {username} created successfully.")
    else:
        print(f"- Failed to create user: {response.text}")

def create_patient(name, nif, email, phone, address, role, token):
    print(f"Creating patient {name}")
    headers = {"Authorization": f"Bearer {token}"}
    data = {
        "name": name,
        "nif": nif,
        "email": email,
        "phone": phone,
        "address": address,
    }
    response = requests.post(f"{BASE_URL}/{role}/patients", json=data, headers=headers)
    if response.status_code == 201:
        print(f"- Patient {name} created successfully.")
    else:
        print(f"- Failed to create patient: {response.text}")

def list_patients(role, token, filters=None):
    print(f"Listing patients with {filters}")
    headers = {"Authorization": f"Bearer {token}"}
    params = {}
    if filters:
        if "name" in filters:
            params["name"] = filters["name"]

    response = requests.get(f"{BASE_URL}/{role}/patients", headers=headers, params=params)
    if response.status_code == 200:
        patients = response.json()
        for patient in patients:
            print(f"- {patient['Name']} ({patient['Email']}) (ID: {patient['ID']})")
            return patient['ID']
    else:
        print(f"- Failed to list patients: {response.text}")


def create_diagnosis(patient_id, diagnosis, prescription, start_date, role, token):
    print(f"Creating diagnosis for patient {patient_id}")
    headers = {"Authorization": f"Bearer {token}"}
    data = {
        "patient_id": patient_id,
        "diagnosis": diagnosis,
        "prescription": prescription,
        "start_date": start_date,
        "createdAt": datetime.date.today().isoformat(),
    }
    response = requests.post(f"{BASE_URL}/{role}/diagnoses", json=data, headers=headers)
    if response.status_code == 201:
        print(f"- Diagnosis for patient {patient_id} created successfully.")
    else:
        print(f"- Failed to create diagnosis: {response.text}")

def list_diagnoses(role, token, filters=None):
    print(f"Listing diagnoses with {filters}")
    headers = {"Authorization": f"Bearer {token}"}
    params = {}
    if filters:
        if "name" in filters:
            params["name"] = filters["name"]
        if "start_date" in filters:
            params["start_date"] = filters["start_date"]

    response = requests.get(f"{BASE_URL}/{role}/diagnoses", headers=headers, params=params)
    if response.status_code == 200:
        diagnoses = response.json()
        for diag in diagnoses:
            print(f"- Patient: {diag['PatientID']} Diagnosis: {diag['Diagnosis']} (Prescription: {diag.get('Prescription', 'None')}) (Start date: {diag['StartDate']})")
    else:
        print(f"- Failed to list diagnoses: {response.text}")


def admin_operations():
    print('-'*25 + "Admin operations" + '-'*25)

    admin_token = login("admin", "notsecurepassword")
    return admin_token


def internal_operations(admin_token):
    print('-'*25 + "Internal operations" + '-'*25)
    today = datetime.date.today().isoformat()
    tommorrow = (datetime.date.today() + datetime.timedelta(days=1)).isoformat()
    
    create_user("InternalUser", "password123", "internal", admin_token)
    internal_token = login("InternalUser", "password123")
    if not internal_token: return
    
    create_patient("Frodo Baggins", "123A", "frodo@example.com", "111-111", "Shire", "internal", internal_token)
    create_patient("Samwise Gamgee", "124B", "sam@example.com", "222-222", "Shire", "internal", internal_token)
    id_frodo = list_patients("internal", internal_token, {"name": "Frodo Baggins"})

    create_diagnosis(id_frodo, "Sprained ankle", "Rest and ice", today, "internal", internal_token)
    create_diagnosis(id_frodo, "Mild fever", "Paracetamol", tommorrow, "internal", internal_token)
    list_diagnoses("internal", internal_token, {"name": "Frodo Baggins"})
    list_diagnoses("internal", internal_token, {"start_date": tommorrow})

    return internal_token

def external_operations(internal_token):
    print('-'*25 + "External operations" + '-'*25)
    today = datetime.date.today().isoformat()
    tommorrow = (datetime.date.today() + datetime.timedelta(days=1)).isoformat()

    create_user("ExternalUser", "password123", "external", internal_token)
    external_token = login("ExternalUser", "password123")
    if not external_token: return
    
    create_patient("Aragorn", "125C", "aragorn@example.com", "333-333", "Rivendell", "external", external_token)
    create_patient("Gimli", "126D", "gimli@example.com", "444-444", "Lonely Mountain", "external", external_token)
    id_aragorn = list_patients("external", external_token, {"name": "Aragorn"})
    
    create_diagnosis(id_aragorn, "Fractured hand", "Immobilization and rest", today, "external", external_token)
    create_diagnosis(id_aragorn, "Allergic reaction", "Antihistamines", tommorrow, "external", external_token)
    list_diagnoses("external", external_token, {"name": "Aragorn"})
    list_diagnoses("external", external_token, {"start_date": tommorrow})   

    return external_token

def main():
    
    print('-'*25 + " Starting scenario..." + '-'*25)
    admin_token = admin_operations()
    if not admin_token: return
    internal_token = internal_operations(admin_token)
    if not internal_token: return
    external_token = external_operations(internal_token)
    print('-'*25 + "Scenario finished." + '-'*25)


if __name__ == "__main__":
    main()
