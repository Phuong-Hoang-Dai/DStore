import { type ReactNode } from "react";

const MainLayout = ({ children }: { children: ReactNode }) => {
  return (
    <div className="h-full w-full md:w-9/10 lg:w-3/4 m-auto px-2">
      {children}
    </div>
  );
};

export default MainLayout;
