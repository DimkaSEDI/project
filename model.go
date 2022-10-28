package main

type GeoResponse []struct {
	Name       string `json:"name"`
	LocalNames struct {
		Ka          string `json:"ka"`
		No          string `json:"no"`
		Ar          string `json:"ar"`
		Ja          string `json:"ja"`
		Fa          string `json:"fa"`
		Nn          string `json:"nn"`
		Ug          string `json:"ug"`
		Co          string `json:"co"`
		Fo          string `json:"fo"`
		Os          string `json:"os"`
		Am          string `json:"am"`
		Ku          string `json:"ku"`
		Uk          string `json:"uk"`
		Ro          string `json:"ro"`
		Da          string `json:"da"`
		Bn          string `json:"bn"`
		Ur          string `json:"ur"`
		Iu          string `json:"iu"`
		Ch          string `json:"ch"`
		Sc          string `json:"sc"`
		Tt          string `json:"tt"`
		Lv          string `json:"lv"`
		Ml          string `json:"ml"`
		Uz          string `json:"uz"`
		Mi          string `json:"mi"`
		Kw          string `json:"kw"`
		Ms          string `json:"ms"`
		Eu          string `json:"eu"`
		Hr          string `json:"hr"`
		Es          string `json:"es"`
		Fr          string `json:"fr"`
		Ty          string `json:"ty"`
		Bo          string `json:"bo"`
		Li          string `json:"li"`
		FeatureName string `json:"feature_name"`
		Br          string `json:"br"`
		An          string `json:"an"`
		Lg          string `json:"lg"`
		Be          string `json:"be"`
		Vo          string `json:"vo"`
		Mk          string `json:"mk"`
		Pl          string `json:"pl"`
		Kv          string `json:"kv"`
		Tg          string `json:"tg"`
		Sl          string `json:"sl"`
		Ay          string `json:"ay"`
		Hi          string `json:"hi"`
		He          string `json:"he"`
		Vi          string `json:"vi"`
		Gl          string `json:"gl"`
		Mt          string `json:"mt"`
		Ie          string `json:"ie"`
		Cy          string `json:"cy"`
		Sv          string `json:"sv"`
		Gv          string `json:"gv"`
		Th          string `json:"th"`
		Sh          string `json:"sh"`
		Cv          string `json:"cv"`
		Yi          string `json:"yi"`
		Se          string `json:"se"`
		Hy          string `json:"hy"`
		Et          string `json:"et"`
		Sg          string `json:"sg"`
		St          string `json:"st"`
		Dv          string `json:"dv"`
		Sk          string `json:"sk"`
		Dz          string `json:"dz"`
		Kl          string `json:"kl"`
		Ia          string `json:"ia"`
		Tl          string `json:"tl"`
		Hu          string `json:"hu"`
		Wo          string `json:"wo"`
		Gn          string `json:"gn"`
		Is          string `json:"is"`
		Tk          string `json:"tk"`
		Ab          string `json:"ab"`
		Za          string `json:"za"`
		Ky          string `json:"ky"`
		Wa          string `json:"wa"`
		Su          string `json:"su"`
		Ta          string `json:"ta"`
		Ss          string `json:"ss"`
		Bs          string `json:"bs"`
		Io          string `json:"io"`
		Ce          string `json:"ce"`
		Sw          string `json:"sw"`
		Ps          string `json:"ps"`
		Pt          string `json:"pt"`
		Te          string `json:"te"`
		Lt          string `json:"lt"`
		Gd          string `json:"gd"`
		Nb          string `json:"nb"`
		ASCII       string `json:"ascii"`
		Nl          string `json:"nl"`
		Zh          string `json:"zh"`
		Ln          string `json:"ln"`
		Na          string `json:"na"`
		Cu          string `json:"cu"`
		Bi          string `json:"bi"`
		Kn          string `json:"kn"`
		Az          string `json:"az"`
		Mr          string `json:"mr"`
		Kk          string `json:"kk"`
		Av          string `json:"av"`
		My          string `json:"my"`
		Yo          string `json:"yo"`
		La          string `json:"la"`
		Cs          string `json:"cs"`
		El          string `json:"el"`
		Kg          string `json:"kg"`
		Oc          string `json:"oc"`
		Fi          string `json:"fi"`
		En          string `json:"en"`
		Af          string `json:"af"`
		Jv          string `json:"jv"`
		Ru          string `json:"ru"`
		So          string `json:"so"`
		Ak          string `json:"ak"`
		Ga          string `json:"ga"`
		ID          string `json:"id"`
		It          string `json:"it"`
		Sr          string `json:"sr"`
		Tr          string `json:"tr"`
		Fy          string `json:"fy"`
		Qu          string `json:"qu"`
		Sq          string `json:"sq"`
		De          string `json:"de"`
		Mn          string `json:"mn"`
		Bg          string `json:"bg"`
		Mg          string `json:"mg"`
		Ba          string `json:"ba"`
		Eo          string `json:"eo"`
		Ko          string `json:"ko"`
		Ht          string `json:"ht"`
		Sm          string `json:"sm"`
		Zu          string `json:"zu"`
		Ca          string `json:"ca"`
	} `json:"local_names"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

type Geo struct {
	Name    string
	Country string
	Lon     float64
	Lat     float64
}

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
