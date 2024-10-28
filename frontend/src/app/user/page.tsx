"use client";
import React, { useEffect } from "react";
import AddressForm from "../../components/addressForm/AddressForm";

const UserPage = () => {
  const [user, setUser] = React.useState<any>({});
  const [addresses, setAddresses] = React.useState<any>([]);
  const [userEmail, setUserEmail] = React.useState<string>("");

  useEffect(() => {
    console.log("User Page");
  }, []);

  useEffect(() => {
    user.ID && fetchAddresses();
  }, [user.ID]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const endpoint =
      process.env.NEXT_PUBLIC_ENDPOINT + `/user?email=${userEmail}`;
    const response = await fetch(endpoint, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();
    setUser(data);
  };

  const fetchAddresses = async () => {
    const endpoint =
      process.env.NEXT_PUBLIC_ENDPOINT + `/addresses?userId=${user.ID}`;
    const response = await fetch(endpoint, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();
    setAddresses(data);
  };

  const handleAddressSubmit = async (address: {
    address_line: string;
    city: string;
    state: string;
    postcode: string;
    country: string;
    building: string;
  }) => {
    const payload = {
      user_id: user.ID,
      ...address,
    };
    console.log("New Address Submitted:", payload);
    const endpoint = process.env.NEXT_PUBLIC_ENDPOINT + "/add-address";
    const res = await fetch(endpoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });
    if (res.ok) {
      alert("Address added successfully!");
    } else {
      const responseBody = await res.json();
      alert(responseBody.message);
    }
    fetchAddresses();
  };
  const handleDeleteAddress = async (addressId: string) => {
    const endpoint =
      process.env.NEXT_PUBLIC_ENDPOINT +
      `/delete-address?address_id=${addressId}`;
    const res = await fetch(endpoint, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (res.ok) {
      alert("Address deleted successfully!");
    } else {
      const responseBody = await res.json();
      alert(responseBody.message);
    }
    fetchAddresses();
  };

  return (
    <div className="flex flex-col gap-4 items-center">
      <h1>User Page</h1>
      <form onSubmit={handleSubmit} className="flex flex-col gap-4">
        <h2>Search User</h2>
        <input
          type="text"
          name="email"
          value={userEmail}
          onChange={(e) => setUserEmail(e.target.value)}
          className="border rounded-md border-black px-4 py-2"
        />
        <button
          type="submit"
          className="bg-green-500 text-white px-4 py-2 rounded-md"
        >
          Submit
        </button>
      </form>
      <div className="flex flex-col gap-4 border-2 border-black p-4 rounded-lg mt-12">
        <h2>ID: {user.ID}</h2>
        <h2>Email: {user.email}</h2>
        <h2>First Name: {user.first_name}</h2>
        <h2>Last Name: {user.last_name}</h2>
        <h2>Country Code: {user.country_code}</h2>
        <h2>Phone Number: {user.phone_number}</h2>
      </div>
      <div className="flex flex-col gap-4 items-center">
        <h2>Addresses</h2>
        <p>Number of addresses: {addresses.length}</p>
        {user.ID && (
          <>
            <AddressForm onSubmit={handleAddressSubmit} />
            <div className="flex flex-col gap-4 border-2 border-black p-4 rounded-lg mt-12">
              {addresses.length > 0 &&
                addresses.map((address: any) => (
                  <div
                    key={address.ID}
                    className="flex gap-4 p-4 border border-black items-center"
                  >
                    <h2>ID: {address.ID}</h2>
                    <h2>AddressLine: {address.addressLine}</h2>
                    <h2>City: {address.city}</h2>
                    <h2>State: {address.state}</h2>
                    <h2>Post Code: {address.postcode}</h2>
                    <h2>Country: {address.country}</h2>
                    <h2>Building: {address.building}</h2>
                    <button
                      onClick={() => handleDeleteAddress(address.ID)}
                      className="bg-red-500 text-white px-4 py-2 rounded-md"
                    >
                      Delete
                    </button>
                  </div>
                ))}
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default UserPage;
