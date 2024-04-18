const Footer = () => {
  return (
    <footer className="bg-gray-200 dark:bg-gray-900">
      <div className="max-w-screen-xl mx-auto p-4 md:grid grid-cols-4 gap-4">
        <div>
          <h3 className="font-bold">Help & Contact</h3>
          <ul>
            <li>Your Account</li>
            <li>Your Orders</li>
            <li>Shipping Rates</li>
            <li>Return Policy</li>
            <li>Payment Methods</li>
            <li>Help</li>
            <li>Contact us</li>
          </ul>
        </div>
        <div>
          <h3 className="font-bold">Product Categories</h3>
          <ul>
            <li>Dogs</li>
            <li>Cats</li>
            <li>Birds</li>
            <li>Rabbits</li>
            <li>Fishes</li>
            <li>Reptiles</li>
            <li>Others</li>
          </ul>
        </div>
        <div>
          <h3 className="font-bold">Company</h3>
          <ul>
            <li>About us</li>
            <li>Career</li>
            <li>Press</li>
            <li>Blog</li>
            <li>Partners</li>
          </ul>
        </div>
        <div>
          <h3 className="font-bold">Follow Us</h3>
          <ul>
            <li>Facebook</li>
            <li>Twitter</li>
            <li>Instagram</li>
            <li>LinkedIn</li>
          </ul>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
