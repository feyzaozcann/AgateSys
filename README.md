# AgateSys

**AgateSys** is a backend system designed to manage various aspects of an advertising agency's operations, including clients, staff, campaigns, and advertisements. Built with Go, it uses the Gin framework to make handling HTTP requests smooth and straightforward.

## Features

- **Client Management**: Create, retrieve, update, and delete client information.
- **Staff Management**: Manage staff details, including roles and grades.
- **Campaign Management**: Oversee advertising campaigns, assign managers, and monitor progress.
- **Advertisement Handling**: Manage advertisements associated with campaigns.

## API Endpoints

### Clients
- `GET /clients`: Retrieve all clients.
- `GET /clients/:id`: Retrieve a specific client by ID.
- `POST /clients`: Create a new client.
- `PUT /clients/:id`: Update an existing client's information.
- `DELETE /clients/:id`: Delete a client.

---

### Staff
- `GET /staff`: Retrieve all staff members.
- `GET /staff/:id`: Retrieve a specific staff member by ID.
- `POST /staff`: Add a new staff member.
- `PUT /staff/:id`: Update a staff member's information.
- `DELETE /staff/:id`: Remove a staff member.

---

### Staff Grades
- `GET /grades`: Retrieve all staff grades.
- `POST /grades`: Create a new staff grade.
- `PUT /grades/:id`: Update a staff grade.
- `DELETE /grades/:id`: Remove a staff grade.

---

### Campaigns
- `GET /campaigns`: Retrieve all campaigns.
- `GET /campaigns/:id`: Retrieve a specific campaign by ID.
- `GET /campaigns/client/:clientID`: Retrieve all campaigns for a specific client.
- `POST /campaigns`: Create a new campaign.
- `PUT /campaigns/:id`: Update an existing campaign's details.
- `PUT /campaigns/:id/manager/:managerID`: Assign a manager to a campaign.
- `DELETE /campaigns/:id`: Delete a campaign.

---

### Campaign Managers
- `GET /campaign-manager`: Retrieve all campaign managers.
- `POST /campaign-manager`: Add a new campaign manager.
- `DELETE /campaign-manager/:id`: Delete a campaign manager.

---

### Advertisements
- `GET /adverts`: Retrieve all advertisements.
- `GET /adverts/:id`: Retrieve a specific advertisement by ID.
- `GET /adverts/campaign/:campaignID`: Retrieve all advertisements for a specific campaign.
- `POST /adverts`: Create a new advertisement.
- `PUT /adverts/:id`: Update an existing advertisement.
- `DELETE /adverts/:id`: Delete an advertisement.


## Project Structure

<pre>
AgateSys/
├── cmd/                # Application entry point
├── db/                 # Database migrations and setup
├── handlers/           # HTTP handlers
├── models/             # Data models
├── repositories/       # Data access layer
├── server/             # Server setup and configuration
├── services/           # Business logic
├── .env                # Environment variables (not included in repo)
├── .gitignore          # Git ignore file
├── go.mod              # Go module file
├── go.sum              # Go dependencies
└── README.md           # Project documentation
</pre>
