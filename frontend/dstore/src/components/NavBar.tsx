import MainLayout from "../layouts/MainLayout";
import SignUpLogin from "../features/auth/SignUpLogin";

const NavBar = () => {
  return (
    <div className="bg-[#ecb996] h-15 md:h-20 w-full">
      <MainLayout>
        <div className="flex flex-row h-full w-full justify-between items-center">
          <img src="../public/logo.png" alt="" className="h-full" />
          <SignUpLogin />
        </div>
      </MainLayout>
    </div>
  );
};

export default NavBar;
