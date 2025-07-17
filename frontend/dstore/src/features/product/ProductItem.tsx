import type { Product } from "./product_model";
import UpdateQuantityOrder from "./UpdateQuantityOrder";

const ProductItem = ({
  product,
  index,
}: {
  product: Product;
  index: number;
}) => {
  return (
    <div
      key={product.id}
      className="pb-5 gap-5 flex flex-col overflow-hidden px-1 w-1/2 md:w-1/3 lg:w-1/5 rounded-2xl shadow-lg bg-white"
    >
      <img
        src="https://i.pinimg.com/736x/81/af/d8/81afd8e8dde180332e8be5e0526c0ba1.jpg"
        alt={product.name}
        className="rounded-t-2xl-2xl "
      />
      <div className="px-3 flex flex-col gap-2">
        <h2 className="font-medium uppercase">{product.name}</h2>
        <span className="text-red-600 font-medium">{product.price}đ</span>
        <UpdateQuantityOrder product={product} />
      </div>
    </div>
  );
};

export default ProductItem;
