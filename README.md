# rest_api_for_startup
REST API for project where startups and investors connect with each other. Here you will see how to work with this API.
# endpoints for post requests
1. /regStartup   

The sample json for request:   

{
    "name": "Test Startup",
    "login": "test_startup",
    "password": "test_password",
    "email": "test@example.com",
    "description": "This is a test startup.",
    "logo": "https://example.com/test_startup_logo.jpg",
    "lowestInvestment": 1000,
    "highestInvestment": 10000,
    "region": "Kazakhstan",
    "website": "https://teststartup.com"
}   
The response of API: "data entered successfully"   
2. /regTeam   

The sample json for request:   

[
    {
        "name":"Sayat",
        "role":"programmer",
        "Description":"someone",
        "startup_id":4
    }, 
    {
        "name":"Sayat",
        "role":"programmer",
        "Description":"someone",
        "startup_id":4
    }
]   
The response of API: "all records were saved"  
3. /regAchievements   

The sample json for request:   

[
    {
        "achievement": "Win in Infomatrix",
        "startupID": 4
    },
    {
        "achievement": "Win in Kaspian",
        "startupID": 4
    }
]   

The response of API: "all records were saved" 
# Get requests  
1. sample url request: /get_startups?region=Kazakhstan&category=&lowestTeam=&highestTeam=&lowestInvestment=&highestInvestment=   

The sample json for response:   

[
  {
    "startup_id": "4",
    "name": "Test Startup",
    "login": "test_startup",
    "password": "test_password",
    "email": "test@example.com",
    "description": "This is a test startup.",
    "logo": "https://example.com/test_startup_logo.jpg",
    "lowestInvestment": 1000,
    "highestInvestment": 10000,
    "region": "Kazakhstan",
    "website": "https://teststartup.com",
    "team_size": 2
  }
]    
2. sample url request: /get_startup?id=4  

The sample json for response:   

{
  "startup_id": "4",
  "name": "Test Startup",
  "login": "test_startup",
  "password": "test_password",
  "email": "test@example.com",
  "description": "This is a test startup.",
  "logo": "https://example.com/test_startup_logo.jpg",
  "lowestInvestment": 1000,
  "highestInvestment": 10000,
  "region": "Kazakhstan",
  "website": "https://teststartup.com",
  "team_size": 2,
  "team": [
    {
    "id": 1,
    "name": "Sayat",
    "role": "programmer",
    "description": "someone",
    "startup_id": 4
    },
    {
    "id": 2,
    "name": "Sayat",
    "role": "programmer",
    "description": "someone",
    "startup_id": 4
    }
  ],
  "achievements": [
    {
    "id": 2,
    "achievement": "Win in Infomatrix",
    "startupID": 4
    },
    {
    "id": 3,
    "achievement": "Win in Kaspian",
    "startupID": 4
    }
  ]
}

