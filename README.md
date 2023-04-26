# rest_api_for_startup
REST API for project where startups and investors connect with each other. Here you will see how to work with this API.
# endpoints for post requests
1. /reg_startup  

The sample json for request:   

{
    "name": "название на кириллице",
    "login": "test_startup",
    "password": "test_password",
    "email": "test@example.com",
    "description": "This is a test startup.",
    "logo": "https://example.com/test_startup_logo.jpg",
    "lowestInvestment": 1000,
    "highestInvestment": 10000,
    "region": "Kazakhstan",
    "website": "https://teststartup.com",
    "industry": "IT"
}   
The response of API: "data entered successfully"   
2. /reg_team   

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
3. /reg_achievements   

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
4. /reg_investor   
The sample json for request:   
{
    "name": "John Doe",
    "login": "johndoe",
    "password": "password123",
    "email": "johndoe@example.com",
    "description": "I'm an angel investor interested in fintech startups.",
    "picture": "https://example.com/pictures/johndoe.jpg",
    "region": "San Francisco",
    "website": "https://johndoe.com",
    "investment": 500000,
    "industry": "Fintech"
}   
The response of API: "data entered successfully"   
5. /reg_cases   
The sample json for request:   
[
		{
			"title": "My Startup Case 1",
			"description": "This is a description of my startup case.",
			"investment": 10000,
			"investor_id": 1
		},
		{
			"title": "My Startup Case 2",
			"description": "This is another description of my startup case.",
			"investment": 20000,
			"investor_id": 1
		}
	]   
The response of API: "all records were saved"    
6. /reg_favourite_startup   
The sample json for request:   
{
    "id": 123,
    "investorID": 456
}   
The response of API: "all data was entered successfully"   
7. /reg_favourite_investor   
The sample json for request:   
{
    "id": 123,
    "startupID": 456
}   
The response of API: "all data was entered successfully"
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
# Patch requests
1. /patch_startup   

the sample request:  
{
    "id": 1,
    "name": "Test Startup",
    "login": "test_startup",
    "password": "test_password",
    "email": "test@example.com",
    "logo": "https://example.com/test_startup_logo.jpg",
    "lowestInvestment": 1000,
    "highestInvestment": 10000,
    "region": "Kazakhstan",
    "website": "https://teststartup.com"
}   
The response of API: "the data was successfully updated"   
2. /patch_startup_description   
the sample request:  
{
    "startup_id": 1,
    "description": "some changes"
}   
The response of API: "the patch request was completed successfully"   
3. /patch_team   
the sample request:  
{
    "id": 5,
    "name":"Sayat",
    "role":"programmer",
    "Description":"someone",
    }   
The response of API: "the patch request was completed successfully"
4. /patch_achievement   
the sample request:  
{
    "id": 5,
    "achievement":"some achievement",
    }   
The response of API: "the patch request was completed successfully"
# Delete requests   
for all delete requests you should just send a json request with id of an element.  
Example:   
{
"id": 3,
}   
response: "the delete request completed successfully"
all delete requests:   
1. /delete_startup
2. /delete_achievement
3. /delete_team_member


