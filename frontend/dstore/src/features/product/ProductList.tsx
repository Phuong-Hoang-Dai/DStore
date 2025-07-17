import { fetchProducts } from "./product_service";
import { useEffect, useState } from "react";
import { type Product, mappingProduct } from "./product_model";
import ProductItem from "./ProductItem";

const ProductList = () => {
  const [products, setProducts] = useState<Product[]>([]);
  useEffect(() => {
    fetchProducts().then((data) => {
      setProducts(mappingProduct(data["data"]));
    });
  }, []);

  const renderProducts = products.map((product, i) => (
    <ProductItem product={product} key={product.id} index={i} />
  ));

  return (
    <>
      <div className="flex flex-wrap flex-row ">{renderProducts}</div>
    </>
  );
};

export default ProductList;
