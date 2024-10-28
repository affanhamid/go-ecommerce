"use client";
import React, { useState } from "react";

interface AddressFormProps {
  onSubmit: (address: {
    address_line: string;
    city: string;
    state: string;
    postcode: string;
    country: string;
    building: string;
  }) => void;
}

const AddressForm: React.FC<AddressFormProps> = ({ onSubmit }) => {
  const [addressLine, setAddressLine] = useState<string>("");
  const [city, setCity] = useState<string>("");
  const [state, setState] = useState<string>("");
  const [postCode, setPostCode] = useState<string>("");
  const [country, setCountry] = useState<string>("");
  const [building, setBuilding] = useState<string>("");
  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  const validate = () => {
    const newErrors: { [key: string]: string } = {};

    if (!addressLine) newErrors.addressLine = "Address Line is required";
    if (!city) newErrors.city = "City is required";
    if (!state) newErrors.state = "State is required";
    if (!postCode) newErrors.postCode = "Post Code is required";
    if (!country) newErrors.country = "Country is required";
    if (!building) newErrors.building = "Building is required";

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (validate()) {
      onSubmit({
        address_line: addressLine,
        city,
        state,
        postcode: postCode,
        country,
        building,
      });
      // Clear form after submit
      setAddressLine("");
      setCity("");
      setState("");
      setPostCode("");
      setCountry("");
      setBuilding("");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4">
      <h2>Add Address</h2>
      <input
        type="text"
        name="building"
        placeholder="Building"
        value={building}
        onChange={(e) => setBuilding(e.target.value)}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors.building && <p className="text-red-500">{errors.building}</p>}
      <input
        type="text"
        name="address_line"
        placeholder="Address Line"
        value={addressLine}
        onChange={(e) => setAddressLine(e.target.value)}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors.addressLine && (
        <p className="text-red-500">{errors.addressLine}</p>
      )}

      <input
        type="text"
        name="city"
        placeholder="City"
        value={city}
        onChange={(e) => setCity(e.target.value)}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors.city && <p className="text-red-500">{errors.city}</p>}

      <input
        type="text"
        name="state"
        placeholder="State"
        value={state}
        onChange={(e) => setState(e.target.value)}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors.state && <p className="text-red-500">{errors.state}</p>}

      <input
        type="text"
        name="postcode"
        placeholder="Post Code"
        value={postCode}
        onChange={(e) => setPostCode(e.target.value)}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors.postCode && <p className="text-red-500">{errors.postCode}</p>}

      <input
        type="text"
        name="country"
        placeholder="Country"
        value={country}
        onChange={(e) => setCountry(e.target.value)}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors.country && <p className="text-red-500">{errors.country}</p>}

      <button type="submit">Submit</button>
    </form>
  );
};

export default AddressForm;
