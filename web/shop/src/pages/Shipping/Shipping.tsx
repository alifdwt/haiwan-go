import TransactionBreadcrumb from "@/components/transaction-breadcrumb";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import useAddress from "./hooks/useAddress";
import { City, Province } from "@/interface/Address";

export default function ShippingPage() {
  const {
    provincesData,
    citiesData,
    isLoadingProvince,
    isLoadingCity,
    setProviceId,
  } = useAddress();

  return (
    <div className="max-w-screen-xl mx-auto">
      <TransactionBreadcrumb />
      <div className="my-5 grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="col-span-2 flex flex-col gap-2">
          <div>
            <p>Provinsi</p>
            {isLoadingProvince ? (
              <div>Loading...</div>
            ) : (
              <Select onValueChange={setProviceId}>
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Provinsi" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Provinsi</SelectLabel>
                    {provincesData.data.rajaongkir.results.map(
                      (province: Province) => {
                        return (
                          <SelectItem
                            key={province.province_id}
                            value={province.province_id}
                          >
                            {province.province}
                          </SelectItem>
                        );
                      }
                    )}
                  </SelectGroup>
                </SelectContent>
              </Select>
            )}
          </div>
          <div>
            <p>Kota</p>
            {isLoadingCity ? (
              <div>Loading...</div>
            ) : (
              <Select>
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Kota" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Kota</SelectLabel>
                    {citiesData &&
                      citiesData.rajaongkir.results.map((city: City) => {
                        return (
                          <SelectItem key={city.city_id} value={city.city_id}>
                            {city.type} {city.city_name}
                          </SelectItem>
                        );
                      })}
                  </SelectGroup>
                </SelectContent>
              </Select>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
