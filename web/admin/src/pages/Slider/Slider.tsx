import { useDispatch, useSelector } from "react-redux";
import { SliderClient } from "./components/client";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { myApi } from "@/helpers/api";
import { getSliders } from "@/slices/slider/sliderSlice";

export default function SliderPage() {
  const dispatch = useDispatch();
  const { access_token } = useSelector((state: RootState) => state.user);
  const { sliders } = useSelector((state: RootState) => state.slider);

  const { isLoading } = useQuery({
    queryKey: ["sliders"],
    queryFn: async () => {
      const { data } = await myApi.get("/slider", {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
      dispatch(getSliders(data.data));
      return data.data;
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SliderClient data={sliders} />
      </div>
    </div>
  );
}
