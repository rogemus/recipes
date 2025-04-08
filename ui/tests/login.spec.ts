import { test, expect } from "@playwright/test";

test("looks ok", async ({ page }) => {
  await page.goto("./login");
  await expect(page).toHaveTitle(/Login/);
  await expect(page).toHaveScreenshot();
});

test("user is able to login", async ({ page }) => {
  await page.goto("./login");
  await expect(page).toHaveTitle(/Login/);

  const email = page.getByTestId("EmailField-input");
  const password = page.getByTestId("PasswordField-input");
  const btn = page.getByTestId("FormSubmit");

  await email.fill("tom@example.com");
  await password.fill("pa55word");

  await btn.click();

  // TODO: NEXT does not update path to dashboard after redirect
  // await page.waitForURL("**/dashboard");
  await expect(page).toHaveTitle(/Dashboard/);
});
