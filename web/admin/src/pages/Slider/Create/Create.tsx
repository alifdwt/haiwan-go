import { useParams } from "react-router-dom";
import { SliderForm } from "./components/slider-form";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { myApi } from "@/helpers/api";
import { useSelector } from "react-redux";

export default function CreateSliderPage() {
  const { sliderId } = useParams();
  const { access_token } = useSelector((state: RootState) => state.user);

  const { data, isLoading } = useQuery({
    queryKey: ["slider", sliderId],
    queryFn: async () => {
      const { data } = await myApi.get(`/slider/${sliderId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
      return data.data;
    },
  });

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        {sliderId === "create" ? (
          <SliderForm initialData={null} />
        ) : isLoading ? (
          <div>Loading...</div>
        ) : (
          <SliderForm initialData={data} />
        )}
      </div>
    </div>
  );
}
