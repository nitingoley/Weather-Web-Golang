Weather Application in Go
This is a simple Go web application that provides weather information for a specified city using the OpenWeatherMap API. The application reads an API key from a configuration file and fetches weather data from the OpenWeatherMap API.

Features
Retrieves real-time weather data by city name.
Displays city name and temperature in Kelvin.
Serves weather data over HTTP in JSON format.
Project Structure
python
Copy code
.
├── main.go                   # Main application code
├── .apiConfig                # Configuration file with API key
├── go.mod                    # Go module file
├── go.sum                    # Go dependencies
└── README.md                 # Project README
Prerequisites
Go 1.16 or later
OpenWeatherMap API Key (create an account at https://openweathermap.org/ to get one)
Getting Started
Clone the repository:

bash
Copy code
git clone https://github.com/your-username/weather-app.git
cd weather-app
Install dependencies:

bash
Copy code
go mod tidy
Set up the configuration:

Create a .apiConfig file in the project root with your OpenWeatherMap API key:

json
Copy code
{
    "OpenWeatherMapApiKey": "your_openweathermap_api_key"
}
