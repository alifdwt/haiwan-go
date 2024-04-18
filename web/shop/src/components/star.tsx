import { Star, StarHalf } from "lucide-react";

const StarRating = ({ rating }: { rating: number }) => {
  const maxStars = 5;
  const fullStars = Math.floor(rating);
  const remainder = rating % 1;
  const hasHalfStar = remainder >= 0.25 && remainder <= 0.75;

  const starElements = [];

  for (let i = 0; i < fullStars; i++) {
    starElements.push(
      <Star fill="yellow" key={i} className="h-5 w-5 text-yellow-300" />
    );
  }

  if (hasHalfStar) {
    starElements.push(
      <div key="half-star" className="relative h-5 w-5">
        <StarHalf
          fill="yellow"
          className="absolute top-0 left-0 h-full w-full text-yellow-300"
        />
      </div>
    );
  }

  const emptyStars = maxStars - starElements.length;
  for (let i = 0; i < emptyStars; i++) {
    starElements.push(
      <Star key={`empty-${i}`} className="h-5 w-5 text-gray-300" />
    );
  }

  return <div className="flex items-center">{starElements}</div>;
};

export default StarRating;
