import type { User } from "./user_model";

const url = import.meta.env.VITE_API_URL + "/user"


export const SignUp = async (user: User) => {
  try{
    const res = await fetch(url, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(user)
    })
    if (!res.ok){
      throw new Error("Network response was not ok");
    }

    const data = await res.json()
    return data
  }catch(error){
    console.error("Failed to fetch products:", error);
    throw error;
  }
}

export const SignIn = async (user: User) => {
  try{
    const res = await fetch(url + "/login", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(user)
    })
    const data = await res.json()
    console.log(data)

    if (!res.ok){
      throw new Error("Network response was not ok" );
    }
    return data
  }catch(error){
    console.error("Failed to fetch products:", error);
    throw error;
  }
}