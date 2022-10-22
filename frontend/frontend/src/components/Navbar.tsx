import * as React from 'react';
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import { Link as RouterLink } from "react-router-dom";
import MenuIcon from "@mui/icons-material/Menu"
import Button from "@mui/material/Button";
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { green } from '@mui/material/colors';
import AppBar from "@mui/material/AppBar";
import { styled, useTheme } from '@mui/material/styles';
import Drawer from '@mui/material/Drawer';
import CssBaseline from '@mui/material/CssBaseline';
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar';
import Divider from '@mui/material/Divider';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';


import HomeRepairServiceIcon from '@mui/icons-material/HomeRepairService';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import PendingActionsIcon from '@mui/icons-material/PendingActions';
import EventNoteIcon from '@mui/icons-material/EventNote';
import HealingIcon from '@mui/icons-material/Healing';
import { MenuItem } from '@mui/material';
import LogoutIcon from '@mui/icons-material/Logout';


const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };
  
  function Navbar() {
    return (
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar>
  
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
            >
              <MenuIcon />
  
            </IconButton>
  
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              ระบบยืมเครื่องมือแพทย์
            </Typography>
  
            <Button color="inherit" onClick={signout}>
              Log out
            </Button>
            
          </Toolbar>
        </AppBar>
      </Box>
    );
  }
  
  export default Navbar;