import { SliderClient } from "./components/client";

export default function SliderPage() {
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SliderClient data={[]} />
      </div>
    </div>
  );
}
