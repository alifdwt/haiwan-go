export interface Province {
  province: string;
  province_id: string;
}

export interface City {
  city_id: string;
  city_name: string;
  postal_code: string;
  province: string;
  province_id: string;
  type: string;
}
