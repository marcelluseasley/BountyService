-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    ID SERIAL PRIMARY KEY,
    Username VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    PasswordHash VARCHAR(255) NOT NULL,
    Role VARCHAR(50) NOT NULL,
    JWT TEXT,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL
);

CREATE TABLE Bounties (
    ID SERIAL PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL,
    Reward VARCHAR(255) NOT NULL,
    Category VARCHAR(255) NOT NULL,
    Status VARCHAR(50) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    Deadline TIMESTAMP NOT NULL,
    CreatorID INT NOT NULL,
    AssigneeID INT,
    FOREIGN KEY (CreatorID) REFERENCES Users(ID),
    FOREIGN KEY (AssigneeID) REFERENCES Users(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
DROP TABLE Bounties;
-- +goose StatementEnd
