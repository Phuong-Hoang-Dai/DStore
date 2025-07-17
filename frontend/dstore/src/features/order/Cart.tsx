import { CiShoppingCart } from "react-icons/ci";
import { selectCart } from "./cartSlice";
import { useSelector } from "react-redux";

const Cart = () => {
  const quantity: number = useSelector(selectCart);
  console.log("Cart quantity:", quantity);
  return (
    <>
      <div className="relative flex text-6xl">
        <CiShoppingCart size={45} />
        <span className="text-white text-xs bg-red-500 px-2 py-1 rounded-4xl absolute right-0 top-0">
          {quantity}
        </span>
      </div>
    </>
  );
};

export default Cart;
