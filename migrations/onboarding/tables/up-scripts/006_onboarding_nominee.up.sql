-- Create Nominee table
CREATE TABLE Nominee (
    nominee_id VARCHAR(50) PRIMARY KEY,
    actor_id VARCHAR(50),
    nominee_relation VARCHAR(20),
    nominee_name VARCHAR(100),
    date_of_birth DATE,
    percentage_share INT,
    gender VARCHAR(10),
    pan_number VARCHAR(20),
    FOREIGN KEY (actor_id) REFERENCES Users(actor_id)
);