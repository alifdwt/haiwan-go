import { myApi } from "@/helpers/api";
import { useQuery } from "@tanstack/react-query";
import { ChangeEvent, useState } from "react";

export default function useAddress() {
  const [provinceId, setProviceId] = useState("");

  const { data: provincesData, isLoading: isLoadingProvince } = useQuery({
    queryKey: ["province"],
    queryFn: async () => {
      const { data } = await myApi.get("/rajaongkir/provinsi");
      return data;
    },
  });

  const { data: citiesData, isLoading: isLoadingCity } = useQuery({
    queryKey: ["city"],
    queryFn: async () => {
      const { data } = await myApi.get(`/rajaongkir/city/${provinceId}`);
      return data.data;
    },
  });

  const handleProvinceChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setProviceId(e.target.value);
  };

  return {
    provincesData,
    citiesData,
    isLoadingProvince,
    isLoadingCity,
    setProviceId,
    provinceId,
    handleProvinceChange,
  };
}
