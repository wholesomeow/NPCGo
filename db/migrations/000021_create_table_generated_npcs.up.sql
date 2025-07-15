CREATE TABLE users.generated_npcs (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,

    name TEXT,
    pronouns TEXT[],

    ocean_id UUID REFERENCES npc_traits.ocean_data(id),
    mice_id UUID REFERENCES npc_traits.mice_data(id),
    rei_id UUID REFERENCES npc_traits.rei_data(id),
    cs_id UUID REFERENCES npc_traits.cs_data(id),
    enneagram_id UUID REFERENCES npc_traits.enneagram_data(id),

    npc_type_id UUID REFERENCES npc_meta.npc_types(id),
    body_type_id UUID REFERENCES npc_meta.body_types(id),
    sex_type_id UUID REFERENCES npc_meta.sex_types(id),
    gender_type_id UUID REFERENCES npc_meta.gender_types(id),
    orientation_type_id UUID REFERENCES npc_meta.orientation_types(id),

    appearance_height_ft INT,
    appearance_height_in INT,
    appearance_weight_lbs INT,
    appearance_height_cm FLOAT,
    appearance_weight_kg FLOAT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users.users(id) ON DELETE CASCADE,

    FOREIGN KEY (ocean_id) REFERENCES npc_traits.ocean_data(id) ON DELETE CASCADE,
    FOREIGN KEY (mice_id) REFERENCES npc_traits.mice_data(id) ON DELETE CASCADE,
    FOREIGN KEY (rei_id) REFERENCES npc_traits.rei_data(id) ON DELETE CASCADE,
    FOREIGN KEY (cs_id) REFERENCES npc_traits.cs_data(id) ON DELETE CASCADE,
    FOREIGN KEY (enneagram_id) REFERENCES npc_traits.enneagram_data(id) ON DELETE CASCADE,

    FOREIGN KEY (npc_type_id) REFERENCES npc_meta.npc_types(id) ON DELETE CASCADE,
    FOREIGN KEY (body_type_id) REFERENCES npc_meta.body_types(id) ON DELETE CASCADE,
    FOREIGN KEY (sex_type_id) REFERENCES npc_meta.sex_types(id) ON DELETE CASCADE,
    FOREIGN KEY (gender_type_id) REFERENCES npc_meta.gender_types(id) ON DELETE CASCADE,
    FOREIGN KEY (orientation_type_id) REFERENCES npc_meta.orientation_types(id) ON DELETE CASCADE
);

-- Enable Row-Level Security on the table
-- ALTER TABLE users.generated_npcs ENABLE ROW LEVEL SECURITY;

-- CREATE POLICY select_own_npcs
-- ON users.generated_npcs
-- FOR SELECT
-- USING (user_id = current_setting('app.current_user_id')::INT);

-- CREATE POLICY insert_own_npcs
-- ON users.generated_npcs
-- FOR INSERT
-- WITH CHECK (user_id = current_setting('app.current_user_id')::INT);

-- CREATE POLICY update_own_npcs
-- ON users.generated_npcs
-- FOR UPDATE
-- USING (user_id = current_setting('app.current_user_id')::INT);

-- CREATE POLICY delete_own_npcs
-- ON users.generated_npcs
-- FOR DELETE
-- USING (user_id = current_setting('app.current_user_id')::INT);

-- Must use "SET app.current_user_id = 'whatever';" before any query