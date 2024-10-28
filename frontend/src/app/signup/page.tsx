"use client";
import React, { useState } from "react";

interface FormData {
  first_name: string;
  last_name: string;
  country_code: string;
  phone_number: string;
  email: string;
  password: string;
  user_type: string;
}

const Input = ({ label, name, type, errors, formData, handleChange }: any) => {
  return (
    <div className="flex flex-col gap-2">
      <label>{label}</label>
      <input
        type={type}
        name={name}
        value={formData[name] || ""}
        onChange={handleChange}
        className="border rounded-md border-black px-4 py-2"
      />
      {errors[name] && <span>{errors[name]}</span>}
    </div>
  );
};

const MyForm: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({
    first_name: "",
    last_name: "",
    country_code: "",
    phone_number: "",
    email: "",
    password: "",
    user_type: "USER",
  });
  const [errors, setErrors] = useState<Partial<FormData>>({});

  const validateForm = () => {
    const newErrors: Partial<FormData> = {};

    if (!formData.first_name) newErrors.first_name = "First Name is required";
    if (!formData.last_name) newErrors.last_name = "Last Name is required";
    if (!formData.country_code)
      newErrors.country_code = "Country Code is required";
    if (!formData.phone_number)
      newErrors.phone_number = "Phone Number is required";
    else if (!/^\d+$/.test(formData.phone_number))
      newErrors.phone_number = "Invalid phone number";
    if (!formData.email) newErrors.email = "Email is required";
    else if (!/\S+@\S+\.\S+/.test(formData.email))
      newErrors.email = "Invalid email format";
    if (!formData.password) newErrors.password = "Password is required";
    else if (formData.password.length < 8)
      newErrors.password = "Password must be at least 8 characters";

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (validateForm()) {
      const endpoint = process.env.NEXT_PUBLIC_ENDPOINT + "/add-user";
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });
      if (response.ok) {
        alert("User added successfully!");
      } else {
        const responseBody = await response.json();
        alert(responseBody.message);
      }
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-col gap-4 w-[450px] mx-auto mt-32"
    >
      <Input
        label="First Name"
        name="first_name"
        type="text"
        errors={errors}
        formData={formData}
        handleChange={handleChange}
      />
      <Input
        label="Last Name"
        name="last_name"
        type="text"
        errors={errors}
        formData={formData}
        handleChange={handleChange}
      />
      <Input
        label="Country Code"
        name="country_code"
        type="text"
        errors={errors}
        formData={formData}
        handleChange={handleChange}
      />
      <Input
        label="Phone Number"
        name="phone_number"
        type="text"
        errors={errors}
        formData={formData}
        handleChange={handleChange}
      />
      <Input
        label="Email"
        name="email"
        type="email"
        errors={errors}
        formData={formData}
        handleChange={handleChange}
      />
      <Input
        label="Password"
        name="password"
        type="password"
        errors={errors}
        formData={formData}
        handleChange={handleChange}
      />

      <button type="submit">Submit</button>
    </form>
  );
};

export default MyForm;
