CREATE TABLE generated_npcs (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,

    name TEXT,
    pronouns TEXT[],

    ocean_id UUID REFERENCES ocean_data(id),
    mice_id UUID REFERENCES mice_data(id),
    rei_id UUID REFERENCES rei_data(id),
    cs_id UUID REFERENCES cs_data(id),
    enneagram_id UUID REFERENCES enneagram_data(id),

    npc_type_id UUID REFERENCES npc_types(id),
    body_type_id UUID REFERENCES body_types(id),
    sex_type_id UUID REFERENCES sex_types(id),
    gender_type_id UUID REFERENCES gender_types(id),
    orientation_type_id UUID REFERENCES orientation_types(id),

    appearance_height_ft INT,
    appearance_height_in INT,
    appearance_weight_lbs INT,
    appearance_height_cm FLOAT,
    appearance_weight_kg FLOAT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users.users(id) ON DELETE CASCADE,

    FOREIGN KEY (ocean_id) REFERENCES ocean_data(id) ON DELETE CASCADE,
    FOREIGN KEY (mice_id) REFERENCES mice_data(id) ON DELETE CASCADE,
    FOREIGN KEY (rei_id) REFERENCES rei_data(id) ON DELETE CASCADE,
    FOREIGN KEY (cs_id) REFERENCES cs_data(id) ON DELETE CASCADE,
    FOREIGN KEY (enneagram_id) REFERENCES enneagram_data(id) ON DELETE CASCADE,

    FOREIGN KEY (npc_type_id) REFERENCES npc_types(id) ON DELETE CASCADE,
    FOREIGN KEY (body_type_id) REFERENCES body_types(id) ON DELETE CASCADE,
    FOREIGN KEY (sex_type_id) REFERENCES sex_types(id) ON DELETE CASCADE,
    FOREIGN KEY (gender_type_id) REFERENCES gender_types(id) ON DELETE CASCADE,
    FOREIGN KEY (orientation_type_id) REFERENCES orientation_types(id) ON DELETE CASCADE
);