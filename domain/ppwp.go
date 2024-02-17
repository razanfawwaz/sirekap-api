package domain

type PpwpData struct {
	Kode            *string `json:"kode"`
	ProvinsiKode    *string `json:"provinsi_kode"`
	ProvinsiNama    *string `json:"provinsi_nama"`
	KabupatenKode   *string `json:"kabupaten_kode"`
	KabupatenNama   *string `json:"kabupaten_nama"`
	KecamatanKode   *string `json:"kecamatan_kode"`
	KecamatanNama   *string `json:"kecamatan_nama"`
	KelurahanKode   *string `json:"kelurahan_kode"`
	KelurahanNama   *string `json:"kelurahan_nama"`
	Tps             *string `json:"tps"`
	SuaraPaslon1    *string `json:"suara_paslon_1"`
	SuaraPaslon2    *string `json:"suara_paslon_2"`
	SuaraPaslon3    *string `json:"suara_paslon_3"`
	CHasil1         *string `json:"chasil_hal_1"`
	CHasil2         *string `json:"chasil_hal_2"`
	CHasil3         *string `json:"chasil_hal_3"`
	SuaraSah        *int    `json:"suara_sah"`
	SuaraTidakSah   *int    `json:"suara_tidak_sah"`
	SuaraTotal      *int    `json:"suara_total"`
	PemilihDptJ     *int    `json:"pemilih_dpt_j"`
	PemilihDptL     *int    `json:"pemilih_dpt_l"`
	PemilihDptP     *int    `json:"pemilih_dpt_p"`
	PenggunaDptJ    *int    `json:"pengguna_dpt_j"`
	PenggunaDptL    *int    `json:"pengguna_dpt_l"`
	PenggunaDptP    *int    `json:"pengguna_dpt_p"`
	PenggunaDptbJ   *int    `json:"pengguna_dptb_j"`
	PenggunaDptbL   *int    `json:"pengguna_dptb_l"`
	PenggunaDptbP   *int    `json:"pengguna_dptb_p"`
	PenggunaTotalJ  *int    `json:"pengguna_total_j"`
	PenggunaTotalL  *int    `json:"pengguna_total_l"`
	PenggunaTotalP  *int    `json:"pengguna_total_p"`
	PenggunaNonDptJ *int    `json:"pengguna_non_dpt_j"`
	PenggunaNonDptL *int    `json:"pengguna_non_dpt_l"`
	PenggunaNonDptP *int    `json:"pengguna_non_dpt_p"`
	Psu             *string `json:"psu"`
	Ts              *string `json:"ts"`
	StatusSuara     bool    `json:"status_suara"`
	StatusAdm       bool    `json:"status_adm"`
	UrlPage         *string `json:"url_page"`
	UrlApi          *string `json:"url_api"`
	CreatedAt       *string `json:"created_at"`
	UpdatedAt       *string `json:"updated_at"`
}

type Daerah struct {
	ProvinsiKode  *string `json:"provinsi_kode,omitempty"`
	ProvinsiNama  *string `json:"provinsi_nama,omitempty"`
	KabupatenKode *string `json:"kabupaten_kode,omitempty"`
	KabupatenNama *string `json:"kabupaten_nama,omitempty"`
	KecamatanKode *string `json:"kecamatan_kode,omitempty"`
	KecamatanNama *string `json:"kecamatan_nama,omitempty"`
	KelurahanKode *string `json:"kelurahan_kode,omitempty"`
	KelurahanNama *string `json:"kelurahan_nama,omitempty"`
	Tps           *string `json:"tps,omitempty"`
}

type HasilPaslon struct {
	Paslon1 *string `json:"paslon_1,omitempty"`
	Paslon2 *string `json:"paslon_2,omitempty"`
	Paslon3 *string `json:"paslon_3,omitempty"`
}

type DataSuara struct {
	SuaraSah      **int `json:"suara_sah"`
	SuaraTidakSah **int `json:"suara_tidak_sah"`
	SuaraTotal    **int `json:"suara_total"`
}

type DataPemilih struct {
	PemilihDptJ     *int `json:"pemilih_dpt_j"`
	PemilihDptL     *int `json:"pemilih_dpt_l"`
	PemilihDptP     *int `json:"pemilih_dpt_p"`
	PenggunaDptJ    *int `json:"pengguna_dpt_j"`
	PenggunaDptL    *int `json:"pengguna_dpt_l"`
	PenggunaDptP    *int `json:"pengguna_dpt_p"`
	PenggunaDptbJ   *int `json:"pengguna_dptb_j"`
	PenggunaDptbL   *int `json:"pengguna_dptb_l"`
	PenggunaDptbP   *int `json:"pengguna_dptb_p"`
	PenggunaTotalJ  *int `json:"pengguna_total_j"`
	PenggunaTotalL  *int `json:"pengguna_total_l"`
	PenggunaTotalP  *int `json:"pengguna_total_p"`
	PenggunaNonDptJ *int `json:"pengguna_non_dpt_j,omitempty"`
	PenggunaNonDptL *int `json:"pengguna_non_dpt_l,omitempty"`
	PenggunaNonDptP *int `json:"pengguna_non_dpt_p,omitempty"`
}

type GambarCHasil struct {
	CHasil1 *string `json:"chasil_hal_1,omitempty"`
	CHasil2 *string `json:"chasil_hal_2,omitempty"`
	CHasil3 *string `json:"chasil_hal_3,omitempty"`
}

type ReportData struct {
	Kode        *string      `json:"kode,omitempty"`
	Daerah      Daerah       `json:"daerah,omitempty"`
	HasilPaslon HasilPaslon  `json:"hasil_paslon,omitempty"`
	DataSuara   DataSuara    `json:"data_suara,omitempty"`
	DataPemilih DataPemilih  `json:"data_pemilih,omitempty"`
	CHasil      GambarCHasil `json:"gambar_chasil,omitempty"`
	UrlPage     *string      `json:"url_page,omitempty"`
	UrlApi      *string      `json:"url_api,omitempty"`
	Psu         *string      `json:"psu,omitempty"`
	Ts          *string      `json:"ts,omitempty"`
	StatusAdm   bool         `json:"status_adm,omitempty"`
	StatusSuara bool         `json:"status_suara,omitempty"`
}
