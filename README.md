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
1. sample url request: /startups?region=New%20York&category=Technology&lowestTeam=5&highestTeam=10&lowestInvestment=10000&highestInvestment=50000&sort=ascending   

The sample json for response:   

[
{
"startup_id": "1",
"name": "ABC Technologies",
"login": "abc_tech",
"password": "abc123",
"email": "abc_tech@gmail.com",
"description": "A technology startup focused on developing software solutions for businesses.",
"logo": "https://example.com/logo.png",
"lowestInvestment": 10000,
"highestInvestment": 50000,
"region": "New York",
"website": "https://abc-tech.com",
"team_size": 8,
"industry": "Technology"
},
{
"startup_id": "2",
"name": "XYZ Innovations",
"login": "xyz_innovations",
"password": "xyz123",
"email": "xyz_innovations@gmail.com",
"description": "An innovative startup focused on creating new products and technologies.",
"logo": "https://example.com/logo.png",
"lowestInvestment": 20000,
"highestInvestment": 80000,
"region": "New York",
"website": "https://xyz-innovations.com",
"team_size": 6,
"industry": "Technology"
}
]    
2. sample url request: /get_startup?id=1   

The sample json for response:   
{
    "startup_id": 1,
    "name": "Startup1",
    "login": "startup1",
    "password": "password",
    "email": "startup1@example.com",
    "description": "A description of Startup1",
    "logo": "https://example.com/startup1-logo.jpg",
    "lowestInvestment": 10000,
    "highestInvestment": 50000,
    "region": "North America",
    "website": "https://startup1.com",
    "team_size": 3,
    "industry": "Technology",
    "team": [
        {
            "id": 1,
            "name": "John Doe",
            "role": "CTO",
            "description": "A description of John Doe",
            "startup_id": 1
        },
        {
            "id": 2,
            "name": "Jane Smith",
            "role": "CEO",
            "description": "A description of Jane Smith",
            "startup_id": 1
        },
        {
            "id": 3,
            "name": "Bob Johnson",
            "role": "COO",
            "description": "A description of Bob Johnson",
            "startup_id": 1
        }
    ],
    "favourites": [
        {
            "id": 1,
            "investorID": 1,
            "startupID": 1
        },
        {
            "id": 2,
            "investorID": 3,
            "startupID": 1
        }
    ],
    "achievements": [
        {
            "id": 1,
            "achievement": "Award 1",
            "startupID": 1
        },
        {
            "id": 2,
            "achievement": "Award 2",
            "startupID": 1
        }
    ]
}   
3. sample url request: /investors?region=California&lowestInvestment=100000&highestInvestment=500000&sort=ascending   

The sample json for response:   
[
{
"investor_id": "1",
"name": "John Doe",
"email": "johndoe@gmail.com",
"description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
"picture": "https://example.com/johndoe.jpg",
"region": "California",
"website": "https://example.com",
"investment": 250000,
"industry": "Technology"
},
{
"investor_id": "2",
"name": "Jane Doe",
"email": "janedoe@gmail.com",
"description": "Nullam eget sapien nec nunc viverra tempus sit amet ac turpis.",
"picture": "https://example.com/janedoe.jpg",
"region": "California",
"website": "https://example.com",
"investment": 400000,
"industry": "Healthcare"
}
]  
4. sample url request: /nvestor?id=1   

The sample json for response:   
{
  "investor_id": 1,
  "name": "John Smith",
  "login": "jsmith",
  "password": "password123",
  "email": "jsmith@example.com",
  "picture": "https://example.com/jsmith.jpg",
  "region": "North America",
  "description": "I'm a seasoned investor with experience in tech startups.",
  "website": "https://example.com/jsmith",
  "investment": 1000000,
  "industry": "Technology",
  "favourites": [
    {
      "id": 1,
      "startup_id": 2,
      "investorID": 1
    },
    {
      "id": 2,
      "startup_id": 5,
      "investorID": 1
    }
  ],
  "cases": [
    {
      "id": 1,
      "title": "Investment in Startup X",
      "description": "I invested $500,000 in Startup X and helped them reach their Series A round.",
      "investment": 500000,
      "investorID": 1
    },
    {
      "id": 2,
      "title": "Investment in Startup Y",
      "description": "I led a $1 million seed round for Startup Y and helped them hire a team of engineers.",
      "investment": 1000000,
      "investorID": 1
    }
  ]
}   
# Patch requests
1. /patch_startup   

the sample request:  
{
  "id": 123,
  "name": "New Startup Name",
  "login": "newlogin",
  "password": "newpassword",
  "email": "newemail@example.com",
  "logo": "https://example.com/newlogo.png",
  "lowestInvestment": 5000,
  "highestInvestment": 10000,
  "region": "New Region",
  "website": "https://newexample.com",
  "industry": "New Industry"
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
4. /patch_investor   
the sample request:  
{
	"id": 1,
	"name": "New Startup Name",
	"login": "newstartuplogin",
	"password": "newstartuppassword",
	"email": "newstartupemail@example.com",
	"logo": "http://example.com/newstartuplogo.png",
	"lowestInvestment": 5000,
	"highestInvestment": 10000,
	"region": "New Startup Region",
	"website": "http://newstartupwebsite.com",
	"industry": "New Startup Industry"
}   
The response of API: "the data was successfully updated"   
5. /patch_startup_achievement   
the sample request:  
{
    "id": 5,
    "achievement":"some achievement",
    }   
The response of API: "the patch request was completed successfully"
6. /patch_investor_description   
the sample request:  
{
  "investor_id": 1,
  "description": "A passionate investor with experience in the tech industry"
}   
The response of API: "the patch request was completed successfully"   
7. /patch_case   
the sample request:  
{
	"id": 1,
	"title": "New Title",
	"description": "New Description",
	"investment": 50000
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
2. /delete_investor
3. /delete_achievement   
4. /delete_case   
5. /delete_favourite_startup   
6. /delete_favourite_investor   
7. /delete_team_member   
# Authorization requests   
1. /auth_startup   
the sample request:  
{
  "login": "example_login",
  "password": "example_password"
}   
The response of API: {
  "id": 1,
  "login": "example_login",
  "error_status": false
}   
2. /auth_investor   
the sample request:  
{
  "login": "example_login",
  "password": "example_password"
}   
The response of API: {
  "id": 1,
  "login": "example_login",
  "error_status": false
}   




