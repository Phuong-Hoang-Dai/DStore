import SignUp from "./features/auth/FormSignUp";
import SignIn from "./features/auth/FormSignIn";
import NavBar from "./components/NavBar";
import ProductList from "./features/product/ProductList";
import MainLayout from "./layouts/MainLayout";

function App() {
  return (
    <>
      <NavBar />
      <div className="bg-gray-100 h-screen w-full">
        <MainLayout>
          <div className=" mt-5">
            <ProductList />
          </div>
        </MainLayout>
      </div>
    </>
  );
}

export default App;
