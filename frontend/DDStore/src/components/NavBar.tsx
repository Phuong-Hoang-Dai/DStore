import UserInfo from "../features/auth/UserInfo";
import Cart from "../features/order/Cart";
import MainLayout from "../layouts/MainLayout";

const NavBar = () => {
  return (
    <>
      <div className="bg-[#ecb996] h-15 md:h-20 w-full fixed z-1000">
        <MainLayout>
          <div className="flex flex-row h-full w-full justify-between items-center">
            <img src="./logo.png" alt="" className="h-full" />
            <div className="flex flex-row items-center h-full">
              <UserInfo />
              <Cart />
            </div>
          </div>
        </MainLayout>
      </div>
      <div className="bg-[#ecb996] h-15 md:h-20 w-full"></div>
    </>
  );
};

export default NavBar;
