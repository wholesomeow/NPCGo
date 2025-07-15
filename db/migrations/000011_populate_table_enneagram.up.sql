COPY generator.enneagram(id, archetype, keyWords, enneagram_description, center, dominant_emotion, fear, desire, wings_1, wings_2, level_of_development_1, level_of_development_2, level_of_development_3, level_of_development_4, level_of_development_5, level_of_development_6, level_of_development_7, level_of_development_8, level_of_development_9, key_motivations, overview, addictions, growth_recommendations_1, growth_recommendations_2, growth_recommendations_3, growth_recommendations_4, growth_recommendations_5)
FROM '/rawdata/csv/Enneagram_Data.csv'
DELIMITER ','
CSV HEADER;