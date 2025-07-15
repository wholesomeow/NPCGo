COPY generator.cognitive_data_npc(id, category, data_name, data_values, data_description)
FROM '/rawdata/csv/NPC_Cognitive_Data.csv'
DELIMITER ','
CSV HEADER;