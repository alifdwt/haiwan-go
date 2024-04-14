import { SliderForm } from "./components/slider-form";

export default function CreateSliderPage() {
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SliderForm initialData={null} />
      </div>
    </div>
  );
}
