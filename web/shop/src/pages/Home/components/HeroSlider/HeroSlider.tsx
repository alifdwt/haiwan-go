import { SliderCarousel } from "@/components/slider-carousel";
import Slider from "@/interface/Slider";

export default function HeroSlider() {
  const sliders: Slider[] = [
    {
      ID: 1,
      CreatedAt: "2024-04-14T14:39:26.115307+07:00",
      UpdatedAt: "2024-04-14T14:39:26.115307+07:00",
      DeletedAt: null,
      name: "Promo 1",
      image:
        "https://res.cloudinary.com/dxirtmo5t/image/upload/v1713080365/HaiwanGo/promo-1.webp",
    },
    {
      ID: 2,
      CreatedAt: "2024-04-14T14:40:54.071015+07:00",
      UpdatedAt: "2024-04-14T14:40:54.071015+07:00",
      DeletedAt: null,
      name: "Promo 2",
      image:
        "https://res.cloudinary.com/dxirtmo5t/image/upload/v1713080453/HaiwanGo/promo-2.webp",
    },
    {
      ID: 3,
      CreatedAt: "2024-04-18T19:25:26.624721+07:00",
      UpdatedAt: "2024-04-18T19:25:26.624721+07:00",
      DeletedAt: null,
      name: "gpjkwd",
      image: "https://picsum.photos/id/76/1208/302",
    },
    {
      ID: 4,
      CreatedAt: "2024-04-18T19:25:26.627717+07:00",
      UpdatedAt: "2024-04-18T19:25:26.627717+07:00",
      DeletedAt: null,
      name: "inxnbh",
      image: "https://picsum.photos/id/65/1208/302",
    },
  ];

  return (
    <div className="flex items-center justify-center">
      <SliderCarousel sliders={sliders} />
    </div>
  );
}
