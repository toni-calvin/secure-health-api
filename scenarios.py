import datetime
import os
import requests
from termcolor import colored
from dotenv import load_dotenv

load_dotenv()

BASE_URL = "http://localhost:8080"
ADMIN_USERNAME = os.getenv("ADMIN_USERNAME")
ADMIN_PASSWORD = os.getenv("ADMIN_PASSWORD")

def print_boxed_message(message, color="blue"):
    line_length = len(message) + 4
    print(colored("+" + "-" * line_length + "+", color))
    print(colored(f"|  {message}  |", color))
    print(colored("+" + "-" * line_length + "+", color))

def print_colored_message(message, color="white"):
    print(colored(message, color))


def login(username, password):
    print_colored_message(f"Logging in as {username}...", "yellow")
    response = requests.post(f"{BASE_URL}/login", json={"username": username, "password": password})
    if response.status_code == 200:
        token = response.json()["token"]
        print_colored_message(f"- Login successful. Token: {token}", "green")
        return token
    else:
        print_colored_message(f"- Login failed: {response.text}", "red")
        return None

def create_user(username, password, role, token):
    print_colored_message(f"Creating {role} user {username}...", "yellow")
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.post(f"{BASE_URL}/{role}/users", json={"username": username, "password": password}, headers=headers)
    if response.status_code == 200:
        print_colored_message(f"- User {username} created successfully.", "green")
    else:
        print_colored_message(f"- Failed to create user: {response.text}", "red")

def create_patient(name, nif, email, phone, address, role, token):
    print_colored_message(f"Creating patient {name}...", "yellow")
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
        print_colored_message(f"- Patient {name} created successfully.", "green")
    else:
        print_colored_message(f"- Failed to create patient: {response.text}", "red")

def list_patients(role, token, filters=None):
    print_colored_message(f"Listing patients with filters {filters}...", "yellow")
    headers = {"Authorization": f"Bearer {token}"}
    params = {}
    if filters:
        if "name" in filters:
            params["name"] = filters["name"]

    response = requests.get(f"{BASE_URL}/{role}/patients", headers=headers, params=params)
    if response.status_code == 200:
        patients = response.json()
        for patient in patients:
            print_colored_message(f"- {patient['Name']} ({patient['Email']}) (ID: {patient['ID']})")
            return patient['ID']
    else:
        print_colored_message(f"- Failed to list patients: {response.text}", "red")

def create_diagnosis(patient_id, diagnosis, prescription, start_date, role, token):
    print_colored_message(f"Creating diagnosis for patient {patient_id}...", "yellow")
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
        print_colored_message(f"- Diagnosis for patient {patient_id} created successfully.", "green")
    else:
        print_colored_message(f"- Failed to create diagnosis: {response.text}", "red")

def list_diagnoses(role, token, filters=None):
    print_colored_message(f"Listing diagnoses with filters {filters}...", "yellow")
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
            print_colored_message(
                f"- Patient: {diag['PatientID']} Diagnosis: {diag['Diagnosis']} (Prescription: {diag.get('Prescription', 'None')}) (Start date: {diag['StartDate']})",
            )
    else:
        print_colored_message(f"- Failed to list diagnoses: {response.text}", "red")

# Admin Operations
def admin_operations():
    print_boxed_message("Admin Operations", color="magenta")
    admin_token = login(ADMIN_USERNAME, ADMIN_PASSWORD)
    return admin_token

# Internal Operations
def internal_operations(admin_token):
    print_boxed_message("Internal Operations", color="magenta")
    today = datetime.date.today().isoformat()
    tomorrow = (datetime.date.today() + datetime.timedelta(days=1)).isoformat()
    
    create_user("InternalUser", "password123", "internal", admin_token)
    internal_token = login("InternalUser", "password123")
    if not internal_token: return
    
    create_patient("Frodo Baggins", "123A", "frodo@example.com", "111-111", "Shire", "internal", internal_token)
    create_patient("Samwise Gamgee", "124B", "sam@example.com", "222-222", "Shire", "internal", internal_token)
    print_colored_message("Listing patients", "cyan")
    id_frodo = list_patients("internal", internal_token, {"name": "Frodo Baggins"})

    create_diagnosis(id_frodo, "Sprained ankle", "Rest and ice", today, "internal", internal_token)
    create_diagnosis(id_frodo, "Mild fever", "Paracetamol", tomorrow, "internal", internal_token)
    
    list_diagnoses("internal", internal_token, {"name": "Frodo Baggins"})
    list_diagnoses("internal", internal_token, {"start_date": tomorrow})

    return internal_token

# External Operations
def external_operations(internal_token):
    print_boxed_message("External Operations", color="magenta")
    today = datetime.date.today().isoformat()
    tomorrow = (datetime.date.today() + datetime.timedelta(days=1)).isoformat()

    create_user("ExternalUser", "password123", "external", internal_token)
    external_token = login("ExternalUser", "password123")
    if not external_token: return
    
    create_patient("Aragorn", "125C", "aragorn@example.com", "333-333", "Rivendell", "external", external_token)
    create_patient("Gimli", "126D", "gimli@example.com", "444-444", "Lonely Mountain", "external", external_token)
    print_colored_message("Listing patients", "cyan")
    id_aragorn = list_patients("external", external_token, {"name": "Aragorn"})
    
    create_diagnosis(id_aragorn, "Fractured hand", "Immobilization and rest", today, "external", external_token)
    create_diagnosis(id_aragorn, "Allergic reaction", "Antihistamines", tomorrow, "external", external_token)
    
    list_diagnoses("external", external_token, {"name": "Aragorn"})
    list_diagnoses("external", external_token, {"start_date": tomorrow})   

    return external_token

def main():
    print_boxed_message("TopDoctors CLI Test Scenarios", color="magenta")
    admin_token = admin_operations()
    if not admin_token:
        return

    internal_token = internal_operations(admin_token)
    if not internal_token:
        return

    external_operations(internal_token)
    print_boxed_message("All Operations Completed Successfully!", color="magenta")

if __name__ == "__main__":
    main()
