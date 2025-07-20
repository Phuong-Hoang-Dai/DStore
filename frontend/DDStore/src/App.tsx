import Banner from "./components/Banner";
import NavBar from "./components/NavBar";
import ProductList from "./features/product/ProductList";
import MainLayout from "./layouts/MainLayout";

function App() {
  return (
    <>
      <NavBar />
      <div className="bg-gray-100 h-auto w-full pt-5">
        <MainLayout>
          <Banner />
          <div className=" mt-5">
            <ProductList />
          </div>
        </MainLayout>
      </div>
    </>
  );
}

export default App;
