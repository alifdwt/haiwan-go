import Slider from "@/interface/Slider";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "./ui/carousel";
import { useRef } from "react";
import Autoplay from "embla-carousel-autoplay";

export function SliderCarousel({ sliders }: { sliders: Slider[] }) {
  const plugin = useRef(
    Autoplay({
      delay: 5000,
      stopOnInteraction: false,
    })
  );

  return (
    <Carousel
      plugins={[plugin.current]}
      onMouseEnter={plugin.current.stop}
      onMouseLeave={plugin.current.reset}
      opts={{ loop: true }}
    >
      <CarouselContent>
        {sliders.map((slider) => (
          <CarouselItem key={slider.ID}>
            <img
              src={slider.image}
              alt={slider.name}
              className="w-full rounded-lg"
            />
          </CarouselItem>
        ))}
      </CarouselContent>
      <CarouselPrevious className="absolute left-0" />
      <CarouselNext className="absolute right-0" />
    </Carousel>
  );
}
