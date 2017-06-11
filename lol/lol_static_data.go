package lol


import (
	"context"
	"fmt"
	"errors"
	"strings"
)

type StaticData service

type StaticChampion struct{
	Info        *InfoDto    		`json:"info,omitempty"`        //InfoDto - This object contains champion information.
	Enemytips   *[]string   		`json:"enemytips,omitempty"`   // Tips about hero
	Stats		*StatsDto			`json:"stats,omitempty"`       //StatsDto - This object contains champion stats data.
	Name		*string				`json:"name,omitempty"`
	Title		*string 			`json:"title,omitempty"`
	Image		*ImageDto 	 		`json:"image,omitempty"`       //ImageDto - This object contains image data.
	Tags		*[]string 			`json:"tags,omitempty"`
	Partype		*string 		  	`json:"partype,omitempty"`
	Skins		*[]SkinDto 			`json:"skins,omitempty"`       //SkinDto - This object contains champion skin data.
	Passive		*PassiveDto 		`json:"passive,omitempty"`     //PassiveDto - This object contains champion passive data.
	Recommended	*[]RecommendedDto   `json:"recommended,omitempty"` //RecommendedDto - This object contains champion recommended data.
	Allytips	*[]string 			`json:"allytips,omitempty"`
	Key			*string 			`json:"key,omitempty"`
	Lore		*string 			`json:"lore,omitempty"`
	ID			*int 				`json:"id,omitempty"`
	Blurb		*string 			`json:"blurb,omitempty"`
	Spells		*[]ChampionSpellDto `json:"spells,omitempty"`      //ChampionSpellDto - This object contains champion spell data.

}

//InfoDto - This object contains champion information.

type InfoDto struct{
	Difficulty	*int `json:"difficulty,omitempty"`
	Attack	    *int `json:"attack,omitempty"`
	Defense	    *int `json:"defense,omitempty"`
	Magic	    *int `json:"magic,omitempty"`
}

//StatsDto - This object contains champion stats data.

type StatsDto struct{
	Armorperlevel		 *int `json:"armorperlevel,omitempty"`
	Hpperlevel			 *int `json:"hpperlevel,omitempty"`
	Attackdamage		 *int `json:"attackdamage,omitempty"`
	Mpperlevel			 *int `json:"mpperlevel,omitempty"`
	Attackspeedoffset	 *int `json:"attackspeedoffset,omitempty"`
	Armor				 *int `json:"armor,omitempty"`
	Hp					 *int `json:"hp,omitempty"`
	Hpregenperlevel		 *int `json:"hpregenperlevel,omitempty"`
	Spellblock			 *int `json:"spellblock,omitempty"`
	Attackrange			 *int `json:"attackrange,omitempty"`
	Movespeed			 *int `json:"movespeed,omitempty"`
	Attackdamageperlevel *int `json:"attackdamageperlevel,omitempty"`
	Mpregenperlevel		 *int `json:"mpregenperlevel,omitempty"`
	Mp					 *int `json:"mp,omitempty"`
	Spellblockperlevel	 *int `json:"spellblockperlevel,omitempty"`
	Crit				 *int `json:"crit,omitempty"`
	Mpregen				 *int `json:"mpregen,omitempty"`
	Attackspeedperlevel	 *int `json:"attackspeedperlevel,omitempty"`
	Hpregen				 *int `json:"hpregen,omitempty"`
	Critperlevel		 *int `json:"critperlevel,omitempty"`
}

//ImageDto - This object contains image data.

type ImageDto struct{
	Full	*string `json:"full,omitempty"`
	Group	*string `json:"group,omitempty"`
	Sprite	*string `json:"sprite,omitempty"`
	H		*int    `json:"h,omitempty"`
	W		*int    `json:"w,omitempty"`
	Y		*int    `json:"y,omitempty"`
	X		*int    `json:"x,omitempty"`
}

//SkinDto - This object contains champion skin data.

type SkinDto struct{
	Num	    *int    `json:"num,omitempty"`
	Name	*string `json:"name,omitempty"`
	ID	    *int    `json:"id,omitempty"`
}

//PassiveDto - This object contains champion passive data.

type PassiveDto struct{
	Image					*ImageDto `json:"image,omitempty"` 					//ImageDto - This object contains image data.
	SanitizedDescription	*string   `json:"sanitizedDescription,omitempty"`
	Name					*string   `json:"name,omitempty"`
	Description				*string   `json:"description,omitempty"`
}

//RecommendedDto - This object contains champion recommended data.

type RecommendedDto struct{
	Map		 *string     `json:"map,omitempty"`
	Blocks	 *[]BlockDto `json:"blocks,omitempty"`    //BlockDto - This object contains champion recommended block data.
	Champion *string     `json:"champion,omitempty"`
	Title	 *string	 `json:"title,omitempty"`
	Priority *bool		 `json:"priority,omitempty"`
	Mode	 *string	 `json:"mode,omitempty"`
	Type	 *string	 `json:"type,omitempty"`
}

//BlockDto - This object contains champion recommended block data.

type BlockDto struct{
	Items	*[]BlockItemDto `json:"items,omitempty"`   //BlockItemDto - This object contains champion recommended block item data.
	RecMath	*bool 			`json:"recMath,omitempty"`
	Type	*string			`json:"type,omitempty"`
}

//BlockItemDto - This object contains champion recommended block item data.

type BlockItemDto struct{
	Count	*int `json:"count,omitempty"`
	ID		*int `json:"id,omitempty"`
}

//ChampionSpellDto - This object contains champion spell data.

type ChampionSpellDto struct{
	CooldownBurn		 *string		 `json:"cooldownBurn,omitempty"`
	Resource			 *string		 `json:"resource,omitempty"`
	Leveltip			 *LevelTipDto  	 `json:"leveltip,omitempty"`  		     //LevelTipDto - This object contains champion level tip data.
	Vars				 *[]SpellVarsDto `json:"vars,omitempty"`				 //SpellVarsDto - This object contains spell vars data.
	CostType			 *string		 `json:"costType,omitempty"`
	Image				 *ImageDto		 `json:"image,omitempty"`
	SanitizedDescription *string		 `json:"sanitizedDescription,omitempty"`
	SanitizedTooltip	 *string		 `json:"sanitizedTooltip,omitempty"`
	Effect				 *[][]int        `json:"effect,omitempty"`  			 //This field is a List of List of Double.
	Tooltip				 *string		 `json:"tooltip,omitempty"`
	Maxrank				 *int			 `json:"maxrank,omitempty"`
	CostBurn			 *string		 `json:"costBurn,omitempty"`
	RangeBurn			 *string	     `json:"rangeBurn,omitempty"`
	Range				 *[]string		 `json:"range,omitempty"`				 //This field is either a List of Integer or the String 'self' for spells that target one's own champion.
	Cooldown			 *[]int		   	 `json:"cooldown,omitempty"`
	Cost				 *[]int			 `json:"cost,omitempty"`
	Key					 *string		 `json:"key,omitempty"`
	Description			 *string		 `json:"description,omitempty"`
	EffectBurn			 *[]string		 `json:"effectBurn,omitempty"`
	Altimages			 *[]ImageDto	 `json:"altimages,omitempty"`
	Name				 *string		 `json:"name,omitempty"`
}

//LevelTipDto - This object contains champion level tip data.

type LevelTipDto struct{
	Effect	*[]string `json:"effect,omitempty"`
	Label	*[]string `json:"label,omitempty"`
}

//SpellVarsDto - This object contains spell vars data.

type SpellVarsDto struct{
	RanksWith	*string `json:"ranksWith,omitempty"`
	Dyn		 	*string `json:"dyn,omitempty"`
	Link		*string `json:"link,omitempty"`
	Coeff		*[]int  `json:"coeff,omitempty"`
	Key			*string `json:"key,omitempty"`
}

func (s *SummonerService) GetChampion(ctx context.Context, id int, tags string, locale string) (*StaticChampion, *Response, error) {
		c := fmt.Sprintf("%v/champions/%v", s.client.StaticDataURL,id)

		if tags{
			c = addParamQuery(c,"tags",tags)
		}
		if locale{
			c = addParamQuery(c,"locale",locale)
		}else{
			c = addParamQuery(c,"locale",s.client.Locale)
		}

		req, err := s.client.NewRequest("GET", c, nil)
		if err != nil {
			return nil, nil, err
		}

		uResp := new(StaticChampion)
		resp, err := s.client.Do(ctx, req, uResp)
		if err != nil {
			return nil, resp, err
		}

		return uResp, resp, nil

}


