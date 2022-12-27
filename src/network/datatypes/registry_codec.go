package datatypes

type RegistryCodec struct {
	MinecraftWorldgenBiome struct {
		Value []struct {
			ID      int `json:"id"`
			Element struct {
				Precipitation string  `json:"precipitation"`
				Temperature   float64 `json:"temperature"`
				Downfall      float64 `json:"downfall"`
				Effects       struct {
					WaterColor int `json:"water_color"`
					MoodSound  struct {
						Sound             string  `json:"sound"`
						Offset            float64 `json:"offset"`
						BlockSearchExtent int     `json:"block_search_extent"`
						TickDelay         int     `json:"tick_delay"`
					} `json:"mood_sound"`
					WaterFogColor int `json:"water_fog_color"`
					FogColor      int `json:"fog_color"`
					SkyColor      int `json:"sky_color"`
				} `json:"effects"`
			} `json:"element"`
			Name string `json:"name"`
		} `json:"value"`
		Type string `json:"type"`
	} `json:"minecraft:worldgen/biome"`
	MinecraftDimensionType struct {
		Value []struct {
			Name    string `json:"name"`
			Element struct {
				Ultrawarm                   int     `json:"ultrawarm"`
				LogicalHeight               int     `json:"logical_height"`
				Infiniburn                  string  `json:"infiniburn"`
				PiglinSafe                  int     `json:"piglin_safe"`
				AmbientLight                float64 `json:"ambient_light"`
				HasSkylight                 int     `json:"has_skylight"`
				Effects                     string  `json:"effects"`
				HasRaids                    int     `json:"has_raids"`
				MonsterSpawnBlockLightLimit int     `json:"monster_spawn_block_light_limit"`
				RespawnAnchorWorks          int     `json:"respawn_anchor_works"`
				Height                      int     `json:"height"`
				HasCeiling                  int     `json:"has_ceiling"`
				MonsterSpawnLightLevel      struct {
					Value struct {
						MaxInclusive int `json:"max_inclusive"`
						MinInclusive int `json:"min_inclusive"`
					} `json:"value"`
					Type string `json:"type"`
				} `json:"monster_spawn_light_level"`
				Natural         int     `json:"natural"`
				MinY            int     `json:"min_y"`
				CoordinateScale float64 `json:"coordinate_scale"`
				BedWorks        int     `json:"bed_works"`
			} `json:"element"`
			ID int `json:"id"`
		} `json:"value"`
		Type string `json:"type"`
	} `json:"minecraft:dimension_type"`
	MinecraftChatType struct {
		Value []struct {
			Name    string `json:"name"`
			ID      int    `json:"id"`
			Element struct {
				Chat struct {
					Decoration struct {
						Parameters     []string `json:"parameters"`
						TranslationKey string   `json:"translation_key"`
						Style          struct {
						} `json:"style"`
					} `json:"decoration"`
				} `json:"chat"`
				Narration struct {
					Decoration struct {
						Parameters     []string `json:"parameters"`
						TranslationKey string   `json:"translation_key"`
						Style          struct {
						} `json:"style"`
					} `json:"decoration"`
					Priority string `json:"priority"`
				} `json:"narration"`
			} `json:"element"`
		} `json:"value"`
		Type string `json:"type"`
	} `json:"minecraft:chat_type"`
}
