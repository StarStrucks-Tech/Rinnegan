-- Create Users table
CREATE TABLE Phone_Verification_Stage (
    phone_verification_id  VARCHAR PRIMARY KEY,
    actor_id VARCHAR(30),
    phone_number VARCHAR(10),
    otp VARCHAR(6),
    phone_verification_status PHONE_VERIFICATION_STATUS NOT NULL DEFAULT 'PHONE_NOT_INITIATED',
    FOREIGN KEY (actor_id) REFERENCES Users(actor_id)
);