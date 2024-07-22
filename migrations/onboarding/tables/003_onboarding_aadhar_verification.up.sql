-- Create AadharVerification table
CREATE TABLE Aadhar_Verification_Stage (
    aadhar_number VARCHAR(50) PRIMARY KEY,
    actor_id VARCHAR(30),
    linked_phone_number VARCHAR(20),
    aadhar_status AADHAR_VERIFICATION_STATUS,
    aadhar_otp VARCHAR(10),
    FOREIGN KEY (actor_id) REFERENCES Users(actor_id)
);
