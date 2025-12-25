# Railway CLI Setup for Windows

## PowerShell Execution Policy Error

If you're getting an error like:
```
File C:\Users\user\AppData\Roaming\npm\railway.ps1 cannot be loaded. The file is not digitally signed.
```

This is because PowerShell's execution policy is blocking unsigned scripts. Here are several ways to fix it:

## Solution 1: Change Execution Policy (Recommended for Current User)

Run PowerShell as Administrator and execute:

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

This allows scripts downloaded from the internet to run if they're signed, and local scripts to run without signing.

## Solution 2: Bypass Execution Policy for Current Session

If you don't want to change the policy permanently, you can bypass it for the current PowerShell session:

```powershell
Set-ExecutionPolicy -ExecutionPolicy Bypass -Scope Process
```

Then run your Railway command. Note: This only works for the current PowerShell window.

## Solution 3: Run Railway Commands Directly with Bypass

You can run Railway commands with the bypass flag:

```powershell
powershell -ExecutionPolicy Bypass -Command "railway [your-command]"
```

## Solution 4: Use Command Prompt (CMD) Instead

You can use Command Prompt instead of PowerShell:

1. Open Command Prompt (cmd.exe)
2. Navigate to your project directory
3. Run Railway commands directly:
   ```cmd
   railway login
   railway link
   railway up
   ```

## Solution 5: Use Railway Web Dashboard

Instead of using the CLI, you can:
1. Go to https://railway.app
2. Log in to your account
3. Configure your database connection string directly in the dashboard
4. Set environment variables through the web interface

## Installing Railway CLI (if not already installed)

If Railway CLI is not installed:

```bash
npm install -g @railway/cli
```

Or using PowerShell (after fixing execution policy):

```powershell
npm install -g @railway/cli
```

## Getting Your Database Connection String

Once you can access Railway:

1. **Via CLI:**
   ```bash
   railway login
   railway link  # Link your project
   railway variables  # View environment variables
   ```

2. **Via Web Dashboard:**
   - Go to https://railway.app
   - Select your PostgreSQL service
   - Click on "Connect" or "Variables"
   - Copy the `DATABASE_URL` or `POSTGRES_URL`

## Setting Environment Variables Locally

After getting your connection string, create a `.env` file in your project root:

```bash
DATABASE_URL=postgresql://user:password@postgres-production-926b.up.railway.app:5432/railway
```

Or set it in PowerShell:

```powershell
$env:DATABASE_URL = "postgresql://user:password@postgres-production-926b.up.railway.app:5432/railway"
```

## Quick Fix Script

You can create a PowerShell script to quickly set the execution policy:

1. Open PowerShell as Administrator
2. Run:
   ```powershell
   Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
   ```
3. Type `Y` when prompted
4. Close and reopen PowerShell

## Verification

After fixing the execution policy, verify Railway CLI works:

```powershell
railway --version
```

If this works, you're all set!

