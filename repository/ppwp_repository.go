package repository

import (
	"database/sql"
	"github.com/razanfawwaz/sirekap-api/domain"
	"log"
)

type PpwpRepository struct {
	db *sql.DB
}

func NewPpwpRepository(db *sql.DB) *PpwpRepository {
	return &PpwpRepository{db}
}

func (r *PpwpRepository) GetAllData(limit int, offset int) ([]domain.ReportData, error) {
	query := `SELECT kode, provinsi_kode, provinsi_nama, kabupaten_kota_kode, kabupaten_kota_nama, kecamatan_kode, kecamatan_nama, kelurahan_desa_kode, kelurahan_desa_nama, tps, suara_paslon_1, suara_paslon_2, suara_paslon_3, suara_sah, suara_tidak_sah, suara_total, pemilih_dpt_j, pemilih_dpt_l, pemilih_dpt_p, pengguna_dpt_j, pengguna_dpt_l, pengguna_dpt_p, pengguna_dptb_j, pengguna_dptb_l, pengguna_dptb_p, pengguna_total_j, pengguna_total_l, pengguna_total_p, pengguna_non_dpt_j, pengguna_non_dpt_l, pengguna_non_dpt_p, chasil_hal_1,chasil_hal_2,chasil_hal_3, psu, ts, status_suara, status_adm, url_page,url_api  FROM ppwp_tps WHERE suara_paslon_1 is not null LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []domain.ReportData

	ppwp := &domain.ReportData{}

	for rows.Next() {
		err := rows.Scan(
			&ppwp.Kode,
			&ppwp.Daerah.ProvinsiKode,
			&ppwp.Daerah.ProvinsiNama,
			&ppwp.Daerah.KabupatenKode,
			&ppwp.Daerah.KabupatenNama,
			&ppwp.Daerah.KecamatanKode,
			&ppwp.Daerah.KecamatanNama,
			&ppwp.Daerah.KelurahanKode,
			&ppwp.Daerah.KelurahanNama,
			&ppwp.Daerah.Tps,
			&ppwp.HasilPaslon.Paslon1,
			&ppwp.HasilPaslon.Paslon2,
			&ppwp.HasilPaslon.Paslon3,
			&ppwp.DataSuara.SuaraSah,
			&ppwp.DataSuara.SuaraTidakSah,
			&ppwp.DataSuara.SuaraTotal,
			&ppwp.DataPemilih.PemilihDptJ,
			&ppwp.DataPemilih.PemilihDptL,
			&ppwp.DataPemilih.PemilihDptP,
			&ppwp.DataPemilih.PenggunaDptJ,
			&ppwp.DataPemilih.PenggunaDptL,
			&ppwp.DataPemilih.PenggunaDptP,
			&ppwp.DataPemilih.PenggunaDptbJ,
			&ppwp.DataPemilih.PenggunaDptbL,
			&ppwp.DataPemilih.PenggunaDptbP,
			&ppwp.DataPemilih.PenggunaTotalJ,
			&ppwp.DataPemilih.PenggunaTotalL,
			&ppwp.DataPemilih.PenggunaTotalP,
			&ppwp.DataPemilih.PenggunaNonDptJ,
			&ppwp.DataPemilih.PenggunaNonDptL,
			&ppwp.DataPemilih.PenggunaNonDptP,
			&ppwp.CHasil.CHasil1,
			&ppwp.CHasil.CHasil2,
			&ppwp.CHasil.CHasil3,
			&ppwp.Psu,
			&ppwp.Ts,
			&ppwp.StatusSuara,
			&ppwp.StatusAdm,
			&ppwp.UrlPage,
			&ppwp.UrlApi,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		result = append(result, *ppwp)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
