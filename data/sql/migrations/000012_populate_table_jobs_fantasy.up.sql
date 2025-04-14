COPY jobs_fantasy(id, category, job_name, alt_name, job_description, can_own, min_status_name, max_status_name, min_status_level, max_status_level)
FROM '/rawdata/csv/NPC_Job_List_Fantasy.csv'
DELIMITER ','
CSV HEADER;