import Categories from "./components/Categories/Categories";
import HeroSlider from "./components/HeroSlider/HeroSlider";
import Trending from "./components/Trending/Trending";

export default function Home() {
  return (
    <div className="space-y-4">
      <HeroSlider />
      <Categories />
      <Trending />
    </div>
  );
}
