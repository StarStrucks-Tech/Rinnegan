
-- Create ExtraDetailsCollection table
CREATE TABLE Extra_Details_Collection_Stage (
    actor_id VARCHAR(50) PRIMARY KEY,
    employment_type EMPLOYMENT_TYPE,
    annual_income INT,
    father_name VARCHAR(100),
    mother_name VARCHAR(100),
    nominee_id VARCHAR(50),
    liveness_otp VARCHAR(10),
    liveness_status LIVENESS_CHECK_STATUS,
    FOREIGN KEY (actor_id) REFERENCES Users(actor_id)
);